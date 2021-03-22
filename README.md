# todo_list_service
TODO List Service

## Main features:
 - create user
 - create tasks
 - get tasks
 - add comment to task
 - remove comment
 
Authentication not implemented. Instead of token, use user ID.

## Quick start
   - Run PostgreSQL DB and put connection credentials into environment variables. (Example - example.env)
   - run migrations with ```make migrate-up```
   - run service with ```make run-server```
   
## CI / CD
 - Partially implemented. Main idea to use separate repository with generated K8s files only (GitOps style).
 - First part of implementation (deployment generation and candidate branch push) can be found in ./.ci/ directory
 - cloudbuild.yaml is suitable for GCP Cloud Build service



