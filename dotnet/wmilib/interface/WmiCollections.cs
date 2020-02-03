// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.


namespace Microsoft.wmi.DotNet.Interface
{
    using System;
    using System.Collections.Generic;
    using System.Linq;
    using System.Text;
    using System.Collections.ObjectModel;
    using System.Collections;
    using System.Linq.Expressions;

    public class WmiPropertyCollection : KeyedCollection<string, IWmiProperty>
    {
        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="property">The property.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(IWmiProperty property)
        {
            return property.Name;
        }


        /// <summary>
        /// Gets the element with the specified key.
        /// </summary>
        /// <param name="propertyName">Name of the property.</param>
        /// <returns></returns>
        public new IWmiProperty this[string propertyName]
        {
            get
            {
                return base[propertyName];
            }
        }
    }

    public class WmiQualifierCollection : KeyedCollection<string, IWmiQualifier>
    {
        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="qualifier">The qualifier.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(IWmiQualifier qualifier)
        {
            return qualifier.Name;
        }

        /// <summary>
        /// Gets the element with the specified key.
        /// </summary>
        /// <param name="propertyName">Name of the property.</param>
        /// <returns></returns>
        public new IWmiQualifier this[string qualifierName]
        {
            get
            {
                return base[qualifierName];
            }
        }
    }

    /// <summary>
    /// WmiMethodParameterCollection
    /// </summary>
    public class WmiMethodParameterCollection :
        DisposableKeyedCollection<string, WmiMethodParameter>
    {
        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="methodParam">The method param.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(WmiMethodParameter methodParam)
        {
            return methodParam.Name;
        }

        /// <summary>
        /// Adds the specified name.
        /// </summary>
        /// <param name="Name">The name.</param>
        /// <param name="Value">The value.</param>
        public void Add(string Name, object Value)
        {
            this.Add(new WmiMethodParameter(Name, Value));
        }

        public WmiMethodParameter Pop(string Name)
        {
            WmiMethodParameter item = this[Name];
            this.Remove(Name);
            return item;
        }

        public T GetValue<T>(string paramName)
        {
            T value = default(T);

            if (this.Contains(paramName))
            {
                WmiMethodParameter item = this[paramName];
                if (item.Value != null)
                    value = (T)Activator.CreateInstance(typeof(T), item.Value);
            }
            return value;
        }

        public T[] GetValueArray<T>(string paramName)
        {

            if (this.Contains(paramName))
            {
                WmiMethodParameter item = this[paramName];
                if (item.Value != null)
                {
                    if (item.Value is Array)
                    {
                        Object[] Value = item.Value as Object[];
                        T[] NewValue = new T[Value.Length];

                        for (int i = 0; i < Value.Length; i++)
                        {
                            NewValue[i] = (T)Activator.CreateInstance(typeof(T), Value[i]);
                        }
                        return NewValue;
                    }
                }
            }
            return null;
        }
    }

#if false
    public class WmiInvokeMethodOptions
    {
        public TimeSpan TimeOut;
    } 
#endif

    public class WmiInstanceCollection : WmiInstanceCollection<IWmiInstance> { }
    /// <summary>
    /// WmiInstanceCollection
    /// </summary>
    public class WmiInstanceCollection<T> : DisposableCollection<T> where T : IWmiInstance 
    {
        public WmiInstanceCollection() : base() { }
        public WmiInstanceCollection(List<T> instances) : base(instances) { }

        public WmiInstanceCollection(T[] instances) : base(instances) { }
        /// <summary>
        /// Gets the first instance.
        /// </summary>
        /// <returns></returns>
        public T GetFirstInstance()
        {
            if (this.Count != 0)
            {
                IEnumerator<T> objEnum = this.GetEnumerator();
                while (objEnum.MoveNext())
                {
                    return objEnum.Current;
                }
            }
            return default(T);
        }

        /// <summary>
        /// Gets Array of the embedded instance of the objects in the collection.
        /// </summary>
        /// <value>The embedded instance.</value>
        public string[]
        GetEmbeddedInstances()
        {
            string[] stringArray = new string[Count];
            int i = 0;
            foreach (T item in this)
            {
                stringArray[i++] = item.EmbeddedInstance;
            }
            return stringArray;
        }

        public T[]
        GetWmiInstances()
        {
            return this.ToArray();
        }

        /// <summary>
        /// 
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <returns></returns>
        public T[] GetInstanceArray<Z>()
        {
           return this.ToArray();
        }
    }

    public class WmiKeyValueCollection : KeyedCollection<string, WmiKeyValue>
    {
        /// <summary>
        /// Gets the key for item.
        /// </summary>
        /// <param name="methodParam">The method param.</param>
        /// <returns></returns>
        protected override string GetKeyForItem(WmiKeyValue keyValue)
        {
            return keyValue.Name;
        }

        /// <summary>
        /// Adds the specified name.
        /// </summary>
        /// <param name="Name">The name.</param>
        /// <param name="Value">The value.</param>
        public void Add(string Name, object Value)
        {
            this.Add(new WmiKeyValue(Name, Value));
        }
    }

    public class WmiKeyValue
    {
        public string Name;
        public object Value;

        public WmiKeyValue(string Name, object Value)
        {
            this.Name = Name;
            this.Value = Value;
        }
    }

    /// <summary>
    /// WmiInstanceCollection
    /// </summary>
    public class WmiClassCollection : DisposableCollection<IWmiClass>
    {
        /// <summary>
        /// Gets the first instance.
        /// </summary>
        /// <returns></returns>
        public IWmiClass GetFirstInstance()
        {
            if (this.Count != 0)
            {
                IEnumerator<IWmiClass> objEnum = this.GetEnumerator();
                while (objEnum.MoveNext())
                {
                    return objEnum.Current;
                }
            }
            return null;
        }
    }


    /// <summary>
    /// 
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class StronglyTypedDisposableCollection<T> : IDisposable where T : IDisposable
    {
        private Collection<T> m_data = new Collection<T>();

        /// <summary>
        /// </summary>
        /// <param name="data">.</param>
        public void Add(T data)
        {
            m_data.Add(data);
        }

        /// <summary>
        /// Gets the enumerator.
        /// </summary>
        /// <returns></returns>
        public IEnumerator GetEnumerator()
        {
            return ((IEnumerable)m_data).GetEnumerator();
        }

        /// <summary>
        /// </summary>
        /// <value></value>
        public T this[int index]
        {
            get { return m_data[index]; }
        }

        /// <summary>
        /// Gets the count.
        /// </summary>
        /// <value>The count.</value>
        public int Count
        {
            get { return m_data.Count; }
        }

        /// <summary>
        /// Clears this instance.
        /// </summary>
        public void Clear()
        {
            m_data.Clear();
        }

        /// <summary>
        /// Removes at.
        /// </summary>
        /// <param name="index">The index.</param>
        public void RemoveAt(
            int index)
        {
            m_data.RemoveAt(index);
        }

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                foreach (var item in this.m_data)
                {
                    using (item) { }
                }
                m_Disposed = true;
            }
        }
        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion
    }

    /// <summary>
    /// 
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class DisposableCollection<T> : Collection<T>, IDisposable where T : IDisposable //, IQueryable<T>
    {

        #region Constructor
        public DisposableCollection() : base() { }
        public DisposableCollection(Collection<T> data) : base(data) { }
        public DisposableCollection(List<T> data) : base(data) { }
        public DisposableCollection(T[] data) : base(data) { }

        #endregion

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                foreach (var item in this)
                {
                    using (item) { }
                }
                m_Disposed = true;
            }
        }

        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion

#if false

        #region Private
        Expression m_Expression;
        QueryProvider m_Provider;
        #endregion
        #region IQueryable Members

        public Type ElementType
        {
            get
            {
                this.ToArray();
                return typeof(T);
            }
        }

        public Expression Expression
        {
            get { return m_Expression; }
        }

        public IQueryProvider Provider
        {
            get { return m_Provider; }
        }

        public IEnumerator<T> GetEnumerator()
        {
            return ((IEnumerable<T>)Provider.Execute(Expression)).GetEnumerator();
        }

        IEnumerator IEnumerable.GetEnumerator()
        {
            return ((IEnumerable)Provider.Execute(Expression)).GetEnumerator();
        }

        public class QueryProvider : IQueryProvider
        {
        #region IQueryProvider Members

            public IQueryable<TElement> CreateQuery<TElement>(Expression expression)
            {
                return new DisposableCollection<TElement>(this, expression);
            }

            public IQueryable CreateQuery(Expression expression)
            {
                Type elementType = expression.Type;
                return (IQueryable)Activator.CreateInstance(
                    typeof(DisposableCollection<>).MakeGenericType(elementType),
                    new object[] { this, expression });
            }

            public TResult Execute<TResult>(Expression expression)
            {
                return (TResult)Execute(expression);
            }

            public object Execute(Expression expression)
            {
                return null;
            }

        #endregion

            public string GetQueryText(Expression expression) { return string.Empty; }
        }

        #endregion
#endif
    }

    #region Linq Provider

    #endregion

    /// <summary>
    /// 
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class DisposableList<T> : List<T>, IDisposable where T : IDisposable
    {
        #region Constructor
        public DisposableList() : base() { }
        public DisposableList(T[] data) : base(data) { }
        #endregion

        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                foreach (var item in this)
                {
                    using (item) { }
                }
                m_Disposed = true;
            }
        }

        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion
    }

    /// <summary>
    /// 
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public class DisposableKeyedCollection<Key, T> : KeyedCollection<Key, T>, IDisposable where T : IDisposable
    {
        #region IDisposable Members
        private bool m_Disposed;
        protected virtual void Dispose(bool disposing)
        {
            if (disposing && !m_Disposed)
            {
                foreach (var item in this)
                {
                    using (item) { }
                }
                m_Disposed = true;
            }
        }

        public void Dispose()
        {
            Dispose(true);
            GC.SuppressFinalize(this);
        }

        #endregion

        protected override Key GetKeyForItem(T item)
        {
            throw new NotImplementedException();
        }
    }
}


