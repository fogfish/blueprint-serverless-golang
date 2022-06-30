# Clean up sandbox

The feature integration into `main` is implemented through pull request (no exceptions whatsoever). CI/CD executes automated pull request deployment to the sandbox environment every time new changes are proposed (each commit). The sandbox environment is a disposable deployment dedicated only for pull request validation, which is destroyed by CI/CD when pull request is merged.
