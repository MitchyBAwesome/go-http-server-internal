# CICD Lunch and Learn

1. Simple Hello World App

2. Run Simple Hello World App in Docker:
    - Dockerfile Basics (Imperative set of commands)
        - FROM
        - RUN
        - ADD / COPY
        - CMD

3. Push Simple Hello World App to Source Control
    - Show Git
    - Create Repo
    - Add Remote

**Now We Need Somewhere to Run This**

4. Create an ECS Cluster and a placeholder Service

5. Deploy Simple Hello World App to running ECS Cluster.

6. Do a Deployment.

### It all starts with Source Code
1. Simple Web App

### Creating a CodeCommit Repo

1. Create CodeCommit Repository called "hello-world"
    - Check Git Version 1.7.9.
    - Generate Git credentials.
    - Clone the Repo

### Building Docker Containers and Dockerfile Explained

1. Login to ecr
```$(aws ecr get-login --no-include-email --region ${REGION})```

2. Build docker image
```docker build -t ${REPO} .```

3. Tag docker image
```docker tag ${REPO}:latest ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${REPO}:latest```
4. Push docker image
```docker push ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/${REPO}:lates```

### Existing Repo 

1. Start pushing my source code to GitHub.
```o```

### Creating a CodeCommit Repo

1. Create CodeCommit Repository called "hello-world"
    - Check Git Version 1.7.9.
    - Generate Git credentials.
    - Clone the Repo

### Creating a Pipeline Using CodePipeline

1. Create CodePipeline Repository called "hello-world"
    - Service Role
    - Artifact Location
2. Setup a Source
    - Public GitHub
        - Choose a Repo
        - Choose a Branch
    - AWS CodeCommit
3. Setup a Build Provider

"imagedefinitions.json"