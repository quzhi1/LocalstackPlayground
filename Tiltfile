# -*- mode: Python -*-

load('ext://helm_remote', 'helm_remote')

# Spin up localstack
helm_remote(
  'localstack',
  repo_name="localstack-repo",
  repo_url='https://helm.localstack.cloud',
  values=['values.localstack.yaml'],
)

# Port forwarding
k8s_resource(
  'localstack',
  port_forwards=4566,
  labels=["localstack"],
)