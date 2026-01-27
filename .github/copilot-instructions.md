# GitHub Copilot Instructions for WMI Repository

## Pull Request Guidelines

### Before Marking PR as Ready for Review

When creating or working on a pull request, **DO NOT** mark the PR as ready for review until all of the following conditions are met:

1. **All CI/CD pipelines must pass**: Wait for all Azure Pipeline builds and checks to complete successfully
   - Build pipeline (`.pipelines/build.yaml`)
   - Static analysis pipeline (`.pipelines/static.yaml`)
   
2. **All status checks must be green**: Verify that all required status checks show a passing status before removing the draft status

3. **All automated tests must pass**: Ensure all test suites complete successfully

4. **Code quality checks must pass**: All linting, formatting, and static analysis tools must pass without errors

### Process

1. Create the PR in **draft** mode initially
2. Make all necessary code changes and commit them
3. Wait for all CI/CD pipelines to start and complete
4. Monitor the status checks on the PR page
5. Only after **all status checks are green and passing**, mark the PR as "Ready for review"

### Important Notes

- If any check fails, investigate and fix the issue before marking as ready for review
- Do not bypass failing checks by marking the PR as ready
- If checks are taking too long, wait for them to complete rather than proceeding prematurely
- Use the GitHub Actions/Azure Pipelines UI to monitor build progress and logs

## Repository Structure

This repository contains Go code for WMI (Windows Management Instrumentation) bindings. When making changes:

- Follow Go best practices and idioms
- Ensure code is properly formatted with `go fmt`
- Run tests locally before pushing when possible
- Update documentation if changing public APIs
