// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

using Microsoft.WmiCodeGen.DotNet.Interface;
using System;
using System.Globalization;
using System.Linq;
using System.Threading;

namespace Microsoft.WmiCodeGen.DotNet.Lib
{
    /// <summary>
    /// 
    /// </summary>
    public class OperationFailedException : WmiException
    {
        public OperationFailedException(string message) : base(message) { }

    }

    /// <summary>
    /// This class would better work with CIM_ConcreteJob and above class
    /// </summary>
    public class WmiJob : WmiInstance
    {
        // <summary>
        /// Cache the amount of time the last operation took. Initialized to -2 as a characteristic invalid value.
        /// </summary>

        #region Constructor
        public WmiJob(IWmiInstance job)
            : base(job)
        {
        }

        #endregion

        #region Properties

        /// <summary>
        /// Gets the completion percentage for the task.
        /// </summary>
        /// <value>The percent complete.</value>
        public ushort
        PercentComplete
        {
            get
            {
                RefreshInstance();
                return this.PercentCompleteNoRefresh;
            }
        }

        private ushort PercentCompleteNoRefresh
        {
            get
            {
                return Convert.ToUInt16(GetProperty("PercentComplete"), CultureInfo.InvariantCulture);
            }
        }

        /// <summary>
        /// Gets whether this task has completed.
        /// </summary>
        /// <value>
        /// 	<c>true</c> if this instance successfully completed; otherwise, <c>false</c>.
        /// 	The property will be set to false if the task  has completed with errors or if
        /// 	the task was terminated / Killed or raised an exception.
        /// </value>
        public bool
        IsCompletedSucessfully
        {
            get
            {
                RefreshInstance();
                return JobState == JobState.Completed;
            }
        }

        /// <summary>
        /// Gets whether this task has completed.
        /// </summary>
        /// <value>
        /// 	<c>true</c> if this instance completed; otherwise, <c>false</c>.
        /// 	This value can be true for any of the following state: sucess, terminated, 
        /// 	exception or completed with errors.
        /// </value>
        public bool
        IsComplete
        {
            get
            {
                RefreshInstance();

                JobState state = JobState;

                //A job is considered completed if it is one of these states.  
                return state == JobState.Completed || state == JobState.Terminated ||
                        state == JobState.Killed || state == JobState.Exception ||
                        state == JobState.CompletedWithErrors;
            }
        }

        /// <summary>
        /// Returns the Job State.
        /// </summary>
        /// <value>The state of the job.</value>
        public JobState
        JobState
        {
            get
            {
                return (JobState)(ushort)GetProperty("JobState");
            }
        }

        /// <summary>
        /// Gets the status of this IVMTask.
        /// </summary>
        /// <value>The status.</value>
        public WmiJobStatus
        Status
        {
            get
            {
                return getStatus();
            }
        }

        /// <summary>
        /// Gets the task-specific error code if the task has completed with errors.
        /// </summary>
        /// <value>The error code.</value>
        public int
        ErrorCode
        {
            get
            {
                return int.Parse(GetProperty("ErrorCode").ToString(), CultureInfo.InvariantCulture);
            }
        }

        /// <summary>
        /// Gets the free-form string that contains the error description.
        /// </summary>
        /// <value>The error description.</value>
        public string
        ErrorDescription
        {
            get
            {
                return GetProperty("ErrorDescription").ToString();
            }
        }

        /// <summary>
        /// Returns actual time taken by the operation with a resolution of microseconds.
        /// On failure returns a timespan with a characteristic invalid value of -3 seconds
        /// </summary>
        public TimeSpan
        ElapsedTime
        {
            get
            {
                try
                {
                    //TODO(jutamez): Verify if not using ManagementDateTimeConverter would work.
                    //return ManagementDateTimeConverter.ToTimeSpan(GetProperty("ElapsedTime").ToString());
                    return (TimeSpan)GetProperty("ElapsedTime");
                }
                catch
                {
                    return new TimeSpan(0, 0, -3);
                }
            }

        }

        /// <summary>
        /// Scheduled Start Time for the Task.
        /// </summary>
        public DateTime
        ScheduledStartTime
        {
            get
            {
                RefreshInstance();
                return (DateTime)GetProperty("ScheduledStartTime");
            }
        }

        /// <summary>
        /// Start Time for the Task.
        /// </summary>
        public DateTime
        StartTime
        {
            get
            {
                RefreshInstance();
                return (DateTime)GetProperty("StartTime");
            }
        }

        /// <summary>
        /// Submitted Time for the Task.
        /// </summary>
        public DateTime
        TimeSubmitted
        {
            get
            {
                RefreshInstance();
                return (DateTime)GetProperty("TimeSubmitted");
            }
        }

        /// <summary>
        /// Gets whether this task has been deleted.
        /// </summary>
        public bool
        IsDeleted
        {
            get
            {
                try
                {
                    RefreshInstance();
                }
                catch (WmiException)
                {
#if false
                    IWmiInstance obj = WMIHelper.GetWmiInstance(WmiInstanceManager, "Select * from " + FRWMIClassNames.ConcreteJob + " Where InstanceId='" + InstanceId + "'");
                    if (obj == null) return true;
                    if (wmiJob != null) wmiJob.Dispose();
                    wmiJob = obj; 
#endif
                }
                return false;
            }
        }

        /// <summary>
        /// Concrete job type
        /// </summary>
        public UInt16 TaskType
        {
            get
            {
                RefreshInstance();

                return (UInt16)GetProperty("JobType");
            }
        }

        #endregion

        #region Private Methods
        private WmiJobStatus getStatus()
        {
            JobState state = JobState;
            WmiJobStatus status = WmiJobStatus.Running;

            switch (state)
            {
                case JobState.New:
                    return WmiJobStatus.New;

                case JobState.Starting:
                case JobState.Running:
                case JobState.Suspended:
                    return WmiJobStatus.Suspend;
                case JobState.ShuttingDown:

                    status = WmiJobStatus.Running;
                    break;

                case JobState.Completed:
                    status = WmiJobStatus.CompletedSuccessfully;
                    break;

                case JobState.CompletedWithErrors:
                    status = WmiJobStatus.CompletedWithErrors;
                    break;

                case JobState.Terminated:
                case JobState.Killed:
                case JobState.Exception:
                    status = WmiJobStatus.Canceled;
                    break;

                default:
                    throw new WmiException("Invalid job state " + state);
            }

            return status;
        }
        #endregion
        #region Public Method
        /// <summary>
        /// Wait for the task to reach a given percentage.
        /// </summary>
        /// <param name="taskPercentage">Percentage the task has to reach</param>
        /// <param name="timeout">Timeout for the wait
        /// If it is set to null, wait forever until the task is complete.</param>
        /// <returns>If the task completed.</returns>
        public void
        WaitForTaskPercentage(
            ushort taskPercentage,
            TimeSpan timeout
            )
        {
            DateTime origin = DateTime.Now;

            while (!IsComplete && PercentComplete < taskPercentage)
            {
                if (timeout != null)
                {
                    TimeSpan timeElapsed = DateTime.Now - origin;
                    if (timeElapsed >= timeout)
                    {
                        throw new TimeoutException("Timed out while waiting for " + GetClassName() + " VMTask to reach " + taskPercentage + "%");
                    }
                }
                Thread.Sleep(100);
            }

            if (Status != WmiJobStatus.Running)
            {
                if (JobState == JobState.Exception)
                {
                    throw new WmiException("The task failed at " + PercentComplete + " with error " + ErrorCode + " -- " + ErrorDescription);
                }
                else
                {
                    throw new WmiException("The task failed to reach " + taskPercentage + "% completion with status " + Status.ToString("G"));
                }
            }
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="jobInstance"></param>
        /// <param name="action"></param>
        /// <param name="percentComplete"></param>
        /// <param name="timeout"></param>
        public static void WaitForAction(IWmiInstance jobInstance,
            UserAction action, uint percentComplete, int timeout)
        {
            using (WmiJob job = new WmiJob(jobInstance))
            {
                job.WaitForAction(action, percentComplete, timeout);
            }
        }

        /// <summary>
        /// WaitForAction
        /// </summary>
        /// <param name="task">VMTask</param>
        /// <param name="action">Wait/Cancel/Async</param>
        /// <param name="percentComplete">Percent to Wait for before Return or Cancel</param>
        /// <param name="timeout">Timeout in Seconds. -1 Infinite</param>
        public void WaitForAction(UserAction action, uint percentComplete, int timeout)
        {
            switch (action)
            {
                case UserAction.Wait:
                    WaitForTaskPercentage(percentComplete, timeout);
                    break;
                case UserAction.Cancel:
                    WaitForTaskPercentage(percentComplete, timeout);
                    Cancel(-1);
                    break;
                case UserAction.None:
                case UserAction.Default:
                case UserAction.Async:
                default:
                    //WaitForTaskPercentage(100, timeout);
                    break;
            }
        }
        /// <summary>
        /// 
        /// </summary>
        /// <param name="taskPercentage"></param>
        /// <param name="timeout"></param>
        public void WaitForTaskPercentage(uint taskPercentage, int timeout)
        {
            DateTime origin = DateTime.Now;
            while (!IsComplete && PercentComplete < taskPercentage)
            {
                TimeSpan timeElapsed = DateTime.Now - origin;
                if (timeout > 0 && timeElapsed.Seconds > timeout)
                {
                    throw new TimeoutException();
                }
                Thread.Sleep(500);
            }

            if (taskPercentage >= 100 && !IsCompletedSucessfully)
            {
                if (this["ErrorCode"] == null)
                {
                    throw new Exception(string.Format(CultureInfo.InvariantCulture,
                        "The ErrorCode value for the Wmi Task {0} is not specified", this["Name"]));
                }
                throw new WmiException((ushort)ErrorCode, ErrorDescription);
            }
        }

        /// <summary>
        /// Wait for the task to complete. If the task fails, throws an exception.
        /// </summary>
        /// <param name="timeout">Timeout for the wait
        /// If it is set to null, wait forever until the task is complete.</param>
        public void WaitForSuccessfulCompletion(
            TimeSpan timeout
            )
        {
            WaitForCompletion(timeout);

            if (Status != WmiJobStatus.CompletedSuccessfully)
            {
                if (JobState == JobState.Exception)
                {
                    throw new WmiException("Failed to complete the task with error code " + this.ErrorCode + "; " + this.ErrorDescription);
                }
                else
                {
                    throw new WmiException("Failed to complete the task with error " + JobState.ToString("G"));
                }
            }
        }

        /// <summary>
        /// Wait for the task to complete.
        /// </summary>
        /// <param name="timeout">Timeout for the wait in milliseconds.
        /// If less than or equal to 0, wait forever until the task is complete.</param>
        public void WaitForCompletion(int timeout)
        {
            this.WaitForCompletion(TimeSpan.FromMilliseconds(timeout));
        }

        /// <summary>
        /// Wait for the task to complete.
        /// </summary>
        /// <param name="timeout">Timeout for the operation to complete in. If timeout is less than or equal to TimeSpan.Zero will wait forever.</param>
        public void WaitForCompletion(TimeSpan timeout)
        {
            // We default the percentage change timeout to 5 minutes. If the total
            // operation timeout wanted is less than that we set the change to the total
            // timeout.
            TimeSpan percentageChangeTimeout = TimeSpan.FromMinutes(5);
            if (percentageChangeTimeout > timeout)
            {
                percentageChangeTimeout = timeout;
            }

            this.WaitForCompletion(timeout, percentageChangeTimeout);
        }

        /// <summary>
        /// Will wait for the task to complete.
        /// </summary>
        /// <param name="operationTimeout">The total time for the operation to complete in.
        /// If the timeout is less than or equal to TimeSpan.Zero we will wait forever.</param>
        /// <param name="progressTimeout">The time between progress updates before calling the operation stuck.
        /// IE: If a job doesnt move from 50% to 51% in this timeout we will throw. If this timeout is greater
        /// than or equal to operationTimeout we will only use the operation timeout.</param>
        /// <exception cref="WmiJobOperationTimeoutException">
        /// Throw if operation time exceeds the operation timeout.
        /// </exception>
        /// <exception cref="WmiJobProgressTimeoutException">
        /// Throw if the task fails to make progress within the progress timeout.
        /// </exception>
        public void WaitForCompletion(TimeSpan operationTimeout, TimeSpan progressTimeout)
        {
            if (progressTimeout > operationTimeout)
            {
                progressTimeout = operationTimeout;
            }
            else if (operationTimeout > TimeSpan.Zero && progressTimeout <= TimeSpan.Zero)
            {
                throw new ArgumentOutOfRangeException("progressTimeout", "Must be greater than TimeSpan.Zero in order to monitor progress.");
            }

            DateTime origin = DateTime.Now;
            DateTime lastOperationTime = origin;
            ushort lastPercentCompleted = 0;
            while (!this.IsComplete)
            {
                if (operationTimeout > TimeSpan.Zero)
                {
                    DateTime currentTime = DateTime.Now;

                    // Make sure we havent gone over on the whole operation.
                    TimeSpan totalTimeElapsed = currentTime - origin;
                    if (totalTimeElapsed >= operationTimeout)
                    {
                        throw new WmiJobOperationTimeoutException("Timed out while waiting for:" + GetClassName() + " VMTask completion");
                    }

                    if (this.PercentCompleteNoRefresh > lastPercentCompleted)
                    {
                        lastPercentCompleted = this.PercentCompleteNoRefresh;
                        lastOperationTime = DateTime.Now;
                    }
                    else
                    {
                        // We havent had a change in percentage progress. Check if we are stuck
                        if ((currentTime - lastOperationTime) >= progressTimeout)
                        {
                            throw new WmiJobProgressTimeoutException(
                                string.Format(CultureInfo.InvariantCulture,
                                    "Timed out at: {0} percent, after making no progress for: {1} seconds, while waiting for: {2}, VMTask to complete.",
                                    this.PercentCompleteNoRefresh,
                                    progressTimeout.TotalSeconds,
                                    GetClassName()));
                        }
                    }
                }

                Thread.Sleep(100);
            }
        }

        /// <summary>
        /// Cancels the task within a specified time.
        /// </summary>
        /// <param name="timeout">Maximum time to cancel the task.</param>
        /// <returns>true if the task was successfully cancelled false otherwise. Throws a TimeoutException in case of a timeout</returns>
        public void
        Cancel(TimeSpan timeout)
        {
            WmiMethodParameterCollection parameters = new WmiMethodParameterCollection();
            parameters.Add("RequestedState", (ushort)RequestedJobState.Terminate);

            this.InvokeMethod("RequestStateChange", parameters);

            // Since we dont want the progress to continue after calling cancel we send this as both the
            // operation and progress timeout. That way we will only fail due to the operation timeout.
            this.WaitForCompletion(timeout, timeout);
        }

        /// <summary>
        /// Cancels the task within a specified time.
        /// </summary>
        /// <param name="timeout">Maximum time to cancel the task (in milliseconds).</param>
        /// <returns>True if cancellation was successful, false if the cancellation failed or timed out</returns>
        public void
        Cancel(int timeout)
        {
            this.Cancel(TimeSpan.FromMilliseconds(timeout));
        }

        /// <summary>
        /// Suspend the job task
        /// </summary>
        /// <param name="timeout">time out for waiting for suspend operation</param>
        /// <returns></returns>
        public bool
        Suspend(
            int timeout
            )
        {
            return RequestStateChange(RequestedJobState.Suspend, timeout);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="state"></param>
        /// <param name="timeout">time in milliseconds</param>
        /// <returns></returns>
        public bool
        RequestStateChange(RequestedJobState state, int timeout)
        {
            WmiMethodParameterCollection parameters = new WmiMethodParameterCollection();
            parameters.Add("RequestedState", (ushort)state);
            InvokeMethod("RequestStateChange", parameters);
            try
            {
                WaitForCompletion(TimeSpan.FromMilliseconds(timeout));
            }
            catch (Exception)
            {
                return false;
            }
            return true;
        }

        /// <summary>
        /// Gets the affected objects.
        /// </summary>
        /// <param name="ClassName">Name of the class.</param>
        /// <returns></returns>
        //TODO(jutamez):Validate this method works
        public IWmiInstance[]
        GetAffectedObjects(
            string ClassName)
        {
            //TODO(jutamez):See if you can return instead a WmiInstanceCollection directly
            WmiInstanceCollection collection = GetAssociatedInstances(ClassName);

            IWmiInstance[] objects = new IWmiInstance[collection.Count];

            int i = 0;
            foreach (IWmiInstance obj in collection)
            {
                objects[i] = obj;
            }

            return objects;
        }

        /// <summary>
        /// Gets the one affected object.
        /// </summary>
        /// <param name="ClassName">Name of the class.</param>
        /// <returns></returns>
        public IWmiInstance
        GetOneAffectedObject(
            string ClassName)
        {
            using (WmiInstanceCollection affectedObjects = GetAssociatedInstances(ClassName))
            {
                if (affectedObjects.Count < 1)
                {
                    throw new WmiException("Unexpected affected object count: " + affectedObjects.Count);
                }

                return affectedObjects.First().Clone() as IWmiInstance;
            }
        }

        /// <summary>
        /// Gets the error chain from the task.
        /// </summary>
        /// <returns>An array of VMError, or null if there are no errors</returns>
        public WmiErrorCollection
        GetErrorEx()
        {
            WmiMethodParameterCollection arguments = new WmiMethodParameterCollection();

            using (WmiMethodResult result = InvokeMethod("GetErrorEx", arguments))
            {
                //TODO(jutamez): Verify this line works as expected
                string[] Errors = result.OutParameters["Errors"].Value as string[];

                if (Errors == null || Errors.Length == 0)
                    return null;

                return new WmiErrorCollection(Errors);
            }
        }
        #endregion
    }


    /// <summary>
    /// The possible JobState values returned from this class's WMI JobState property.
    /// </summary>
    public enum JobState
    {
        /// <summary>
        /// New
        /// </summary>
        New = 2,
        /// <summary>
        /// Starting
        /// </summary>
        Starting = 3,
        /// <summary>
        /// Running
        /// </summary>
        Running = 4,
        /// <summary>
        /// Suspended
        /// </summary>
        Suspended = 5,
        /// <summary>
        /// Shutting Down
        /// </summary>
        ShuttingDown = 6,
        /// <summary>
        /// Completed
        /// </summary>
        Completed = 7,
        /// <summary>
        /// Terminated (i.e. cancelled)
        /// </summary>
        Terminated = 8,
        /// <summary>
        /// Killed
        /// </summary>
        Killed = 9,
        /// <summary>
        /// Exception
        /// </summary>
        Exception = 10,
        /// <summary>
        /// CompletedWithErrors
        /// </summary>
        CompletedWithErrors = 32768
    }

    /// <summary>
    /// The possible value to pass to this class's WMI RequestJobStateChange method.
    /// </summary>
    public enum RequestedJobState
    {
        /// <summary>
        /// Start
        /// </summary>
        Start = 2,
        /// <summary>
        /// Suspend
        /// </summary>
        Suspend = 3,
        /// <summary>
        /// Terminate
        /// </summary>
        Terminate = 4,
        /// <summary>
        /// Kill
        /// </summary>
        Kill = 5
    }

    /// <summary>
    /// , enum which can be used for post action after performing an operation
    /// </summary>
    public enum UserAction
    {
        /// <summary>
        /// Ignore this UserAction
        /// </summary>
        None,
        /// <summary>
        /// Will trigger the function to return immediately 
        /// </summary>
        Async,
        /// <summary>
        /// Will trigger the function to return after waiting for X Percentage
        /// </summary>
        Wait,
        /// <summary>
        /// will trigger the function to cancel the task after waiting for X Percentage
        /// </summary>
        Cancel,
        /// <summary>
        /// Defailt is to Wait to 100 % for given timeout
        /// </summary>
        Default,
    }

    public struct JobTimeout
    {
        const int OneSecond = 1 * 1000;
        const int FiveSeconds = 5 * 1000;
        const int OneMinute = 60 * OneSecond;
        const int OneHour = 60 * OneMinute;
        const int Infinite = -1;
    }

    /// <summary>
    /// The status of an IVMTask. 
    /// </summary>
    /// <remarks>
    /// Note that Viridian tasks do not use all of the CIM defined states so we map them to a smaller subset. 
    /// </remarks>
    public enum WmiJobStatus
    {
        /// <summary>
        /// The task is scheduled.
        /// </summary>
        New,
        /// <summary>
        /// The task is currently running.
        /// </summary>
        Running,
        /// <summary>
        /// The task completed successfully.
        /// </summary>
        CompletedSuccessfully,
        /// <summary>
        /// The task was canceled.
        /// </summary>
        Canceled,
        /// <summary>
        /// The task completed with an error.
        /// </summary>
        CompletedWithErrors,
        /// <summary>
        /// The task is suspended but not completed.
        /// </summary>
        Suspend,
        /// <summary>
        /// The task is Killed or do not exist.
        /// </summary>
        Killed
    }

    /// <summary>
    /// Exception used to signify the timeout of a specific operation.
    /// </summary>
    public class WmiJobOperationTimeoutException : TimeoutException
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="WmiJobOperationTimeoutException"/> class.
        /// </summary>
        /// <param name="message">The message to be shown in the exception.</param>
        public WmiJobOperationTimeoutException(string message)
            : base(message)
        {
        }
    }

    /// <summary>
    /// Exception used to signify the timeout of a specific operation because it is no longer making progress.
    /// </summary>
    public class WmiJobProgressTimeoutException : TimeoutException
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="WmiJobProgressTimeoutException"/> class.
        /// </summary>
        /// <param name="message">The message to be shown in the exception.</param>
        public WmiJobProgressTimeoutException(string message)
            : base(message)
        {
        }
    }
}
