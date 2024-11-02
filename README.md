# AI Powered Discord Chatbot

This repository contains an AI powered Discord chatbot written in Golang. The bot connects to a specified Discord server and uses GPT-4o as its backend to generate replies. The bot only replies when tagged in a message.

## Prerequisites

- Go 1.23 or later
- Docker
- Kubernetes cluster
- GitHub Actions

## Environment Variables

The bot requires the following environment variables to be set:

- `DISCORD_TOKEN`: The token for your Discord bot.
- `GPT4O_API_KEY`: The API key for GPT-4o.

## Configuration

The bot reads its configuration from a Kubernetes ConfigMap mounted to the container. Ensure that the ConfigMap is properly set up in your Kubernetes cluster.

## Running the Bot

1. Clone the repository:

   ```sh
   git clone https://github.com/githubnext/workspace-blank.git
   cd workspace-blank
   ```

2. Set the required environment variables:

   ```sh
   export DISCORD_TOKEN=your_discord_token
   export GPT4O_API_KEY=your_gpt4o_api_key
   ```

3. Build and run the bot:

   ```sh
   go build -o bot .
   ./bot
   ```

## Docker

A Dockerfile is provided to build a Docker image for the bot.

### Building the Docker Image

```sh
docker build -t your_docker_image_name .
```

### Running the Docker Container

```sh
docker run -e DISCORD_TOKEN=your_discord_token -e GPT4O_API_KEY=your_gpt4o_api_key your_docker_image_name
```

## GitHub Actions

A GitHub Actions workflow is provided to build and push the Docker image to a local Harbor instance.

### Workflow Configuration

The workflow uses the self-hosted runner `nomsterbuilder` for building the container. The connection information for the container upload is set as secrets in the GitHub repository. The required secrets are:

- `HARBOR_HOST`
- `HARBOR_USER`
- `HARBOR_TOKEN`

### Running the Workflow

The workflow is triggered on every push to the `main` branch. To manually trigger the workflow, push a commit to the `main` branch:

```sh
git add .
git commit -m "Trigger workflow"
git push origin main
```
