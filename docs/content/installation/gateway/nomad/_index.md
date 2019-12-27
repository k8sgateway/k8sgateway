---
title: "Installing Gloo Gateway on HashiCorp Nomad"
menuTitle: "Installing on Nomad"
description: How to install Gloo to run in Gateway Mode on Nomad, with routing example.
weight: 6
---

Gloo Gateway can be used as an Ingress/Gateway for the Nomad platform. This guide walks through the process of installing Gloo on Nomad, using Consul for service discovery/configuration and Vault for secret storage.

[HashiCorp Nomad](https://www.nomadproject.io/) is a popular workload scheduler that can be used in place of, or in combination with Kubernetes as a way of running long-lived processes on a cluster of hosts. Nomad supports native integration with Consul and Vault, making configuration, service discovery, and credential management easy for application developers.

You can see a demonstration of Gloo using Consul, Nomad, and Vault in this YouTube video.

<iframe src="https://www.youtube.com/embed/7Mk5r9P4kb0" frameborder="0" width="560" height="315" allowfullscreen></iframe>

---

## Architecture

<img src="{{% versioned_link_path fromRoot="/img/gloo-architecture-nomad-consul-vault.png" %}}" alt="Gloo Gateway on Nomad Architecture" width="50%">

Gloo Gateway on Nomad uses multiple pieces of software for deployment and functionality.

- **Docker**: The components of Gloo are deployed as containers running discovery, proxy, envoy, and the gateway
- **Nomad**: Nomad is responsible for scheduling, creating, and maintaining containers hosting the Gloo components
- **Consul**: Consul is used to store key/value pairs that represent the configuration of Gloo
- **Vault**: Vault is used to house sensitive data used by Gloo
- **Levant**: Levant is a template and deployment tool for Nomad jobs allowing the use of a variable file with the Nomad job for Gloo deployment
- **Glooctl**: Command line tool for installing and configuring Gloo

## Preparing for Installation

Before proceeding to the installation, you will need to complete some prerequisites.

### Prerequisite Software

Installation on Nomad requires the following:

- [Levant](https://github.com/jrasell/levant) installed on your local machine
- [Docker](https://docs.docker.com/), [Consul](https://www.consul.io), [Vault](https://www.vaultproject.io), and [Nomad](https://www.nomadproject.io/) installed on the target host machine (which can also be your local machine)

### Download the Installation Files

This tutorial uses files stored on the [Gloo GitHub repository](https://github.com/solo-io/gloo).

In order to install Gloo on Nomad, you'll want to clone the repository:

```
git clone --branch v1.1.0 https://github.com/solo-io/gloo
cd gloo/install/nomad
```

The files used for installation live in the `install/nomad` directory.

```bash
├── demo.sh
├── gloo-policy.hcl
├── jobs
│   ├── gloo.nomad
│   └── petstore.nomad
├── launch-consul-vault-nomad-dev.sh
├── README.md
├── Vagrantfile
└── variables
    ├── variables-linux.yaml
    └── variables-mac.yaml
```

The Gloo [Nomad Job](https://www.nomadproject.io/docs/job-specification/index.html) and the Pet Store job are in the `jobs` directory.

{{% notice note %}}
The gloo.nomad job is experimental and designed to be used with a specific Vault + Consul + Nomad setup.
{{% /notice %}}

The [Levant Variables](https://github.com/jrasell/levant) for the Gloo Nomad Job are in the `variables` directory.

Inputs for the job can be tweaked by modifying `variables/variables-*.yaml` files.

---

## Deploying Gloo with Nomad

The scripts and files included in the Gloo repository provide three different options for deployment:

- [Deploy the entire demonstration environment](#run-the-complete-demo) including the Pet Store application using an all-in-one script
- [Use Vagrant to deploy a VM](#running-nomad-using-vagrant) with the supporting services and Gloo Gateway, then deploy the Pet Store application
- [Deploy the supporting services](#running-nomad-consul-and-vault), Gloo Gateway, and Pet Store application separately

### Run the complete Demo

If your environment is set up with Docker, Nomad, Consul, Vault, and Levant, you can simply run `demo.sh` to create a local demo of Gloo routing to the PetStore Nomad. The script will spin up dev instances of Consult, Nomad, and Vault. Then it will use Nomad to deploy the Gloo Gateway and the Pet Store application. Finally, it will create a route on the Gloo Gateway to the Pet Store application.

```bash
./demo.sh
```

After the script completes its setup process, you can test out the routing rule on Gloo by running the following command.

```bash
curl <nomad-host>:8080/
```

If running on macOS or with Vagrant:

```bash
curl localhost:8080/
```

If running on Linux, use the Host IP on the `docker0` interface:

```bash
curl 172.17.0.1:8080/
```

The value returned should be:

```json
[{"id":1,"name":"Dog","status":"available"},{"id":2,"name":"Cat","status":"pending"}]
```

### Running Nomad Using Vagrant 

The provided `Vagrantfile` will run Nomad, Consul, and Vault inside a VM on your local machine. 

First download and install [HashiCorp Vagrant](https://www.vagrantup.com).

Then run the following command:

```bash
vagrant up
```

Ports will be forwarded to your local system, allowing you to access services on the following ports (on `localhost`):

|  service  | port |
| ----- | ---- |  
| nomad | 4646 | 
| consul | 8500 | 
| vault | 8200 | 
| gloo/http | 8080 | 
| gloo/https | 8443 | 
| gloo/admin | 19000 | 

### Running Nomad, Consul, and Vault

If you've installed Nomad/Consul/Vault locally, you can use `launch-consul-vault-nomad-dev.sh` to run them on your local system.

If running locally (without Vagrant) on macOS, you will need to install the [Weave Net Docker Plugin](https://www.weave.works/docs/net/latest/install/plugin/plugin-v2/):

```bash
docker swarm init # if your docker host is not currently a swarm manager
docker plugin install weaveworks/net-plugin:latest_release --grant-all-permissions
docker plugin enable weaveworks/net-plugin:latest_release
docker network create --driver=weaveworks/net-plugin:latest_release --attachable weave
```

If running locally on Linux, you'll need to disable SELinux in order to run the demo (or add permission for docker containers to access `/` on their filesystem):

```bash
sudo setenforce 0
```

Then run the `launch-consul-vault-nomad-dev.sh` script.

```bash
./launch-consul-vault-nomad-dev.sh
```

The script will launch a dev instance of Consul, Vault, and Nomad and then continue to monitor the status of those services in debug mode. You can stop all of the services by hitting `Ctrl-C`.

Once you have finished launching these services, you are now ready to install Gloo on either your [Linux](#installing-gloo-on-nomad-linux) or [macOS](#installing-gloo-on-nomad-mac) system.

---

### Installing Gloo on Nomad

Once you have a base environment set up with Consul, Vault, and Nomad running, you are ready to deploy the Nomad job that creates the necessary containers to run Gloo. The next two sections will guide you on installing Gloo on [Linux](#installing-gloo-on-nomad-linux) or [macOS](#installing-gloo-on-nomad-mac).

### Installing Gloo on Nomad (Linux)

In this step we will deploy Gloo using Levant on a Linux-based system. The assumption is that you are running Consul, Nomad, and Vault either locally or remotely. 

If you are running these services remotely, then you will need to update the `address` and `consul-address` values with your configuration. The default port for Nomad is 4646 and for Consul is 8500. Make sure to give the full address to your Nomad and Consul servers, e.g. https://my.consul.local:8500.

```bash
levant deploy \
    -var-file variables/variables-linux.yaml \
    -address http://<nomad-host>:<nomad-port> \
    -consul-address http://<consul-host>:<consul-port> \
    jobs/gloo.nomad
```

If running locally or with `vagrant`, you can omit the `address` flags from the deployment command:

```bash
levant deploy \
    -var-file variables/variables-linux.yaml \
    jobs/gloo.nomad
```

You can monitor the status of the deployment job by executing the following command:

```bash
nomad job status gloo
```

When the deployment is complete, you are ready to [deploy the Pet Store application](#deploying-a-sample-application) to demonstrate Gloo's capabilities.

### Installing Gloo on Nomad (Mac)

In this step we will deploy Gloo using Levant on a macOS-based system. The assumption is that you are running Consul, Nomad, and Vault locally.

```bash
levant deploy \
    -var-file variables/variables-mac.yaml \
    jobs/gloo.nomad
```

You can monitor the status of the deployment job by executing the following command:

```bash
nomad job status gloo
```

When the deployment is complete, you are ready to [deploy the Pet Store application](#deploying-a-sample-application) to demonstrate Gloo's capabilities.

---

## Deploying a Sample Application

In this step we will deploy a sample application to demonstrate the capabilities of the Gloo Gateway on either your [Linux](#deploy-the-pet-store-on-nomad-linux) or [macOS](#deploy-the-pet-store-on-nomad-mac) system. We're going to deploy the Pet Store application to Nomad using Levant.

### Deploy the Pet Store on Nomad (Linux)

We will deploy the Pet Store application using Levant and Nomad on your local or remote Linux machine.

If you are running these services remotely, then you will need to update the `address` and `consul-address` values with your configuration. The default port for Nomad is 4646 and for Consul is 8500. Make sure to give the full address to your Nomad and Consul servers, e.g. https://my.consul.local:8500.

```bash
levant deploy \
    -var-file variables/variables-linux.yaml \
    -address <nomad-host>:<nomad-port> \
    -consul-address <consul-host>:<consul-port> \
    jobs/petstore.nomad
```

If running locally or with `vagrant`, you can omit the `address` flags from the deployment command:

```bash
levant deploy \
    -var-file variables/variables-linux.yaml \
    jobs/petstore.nomad
```

You can monitor the status of the deployment job by executing the following command:

```bash
nomad job status petstore
```

When the deployment is complete, you are ready to [create a route for the Pet Store application](#create-a-route-to-the-petstore).

### Deploy the Pet Store on Nomad (Mac)

We will deploy the Pet Store application using Levant and Nomad on your local macOS machine. 

```bash
levant deploy \
    -var-file variables/variables-mac.yaml \
    jobs/petstore.nomad
```

You can monitor the status of the deployment job by executing the following command:

```bash
nomad job status petstore
```

When the deployment is complete, you are ready to [create a route for the Pet Store application](#create-a-route-to-the-petstore).

### Create a Route to the PetStore

We can now use `glooctl` to create a route to the Pet Store app we just deployed:

```bash
glooctl add route \
    --path-prefix / \
    --dest-name petstore \
    --prefix-rewrite /api/pets \
    --use-consul
```

```bash
{"level":"info","ts":"2019-08-22T17:15:24.117-0400","caller":"selectionutils/virtual_service.go:100","msg":"Created new default virtual service","virtualService":"virtual_host:<domains:\"*\" > status:<> metadata:<name:\"default\" namespace:\"gloo-system\" > "}
+-----------------+--------------+---------+------+---------+-----------------+--------------------------------+
| VIRTUAL SERVICE | DISPLAY NAME | DOMAINS | SSL  | STATUS  | LISTENERPLUGINS |             ROUTES             |
+-----------------+--------------+---------+------+---------+-----------------+--------------------------------+
| default         |              | *       | none | Pending |                 | / -> gloo-system.petstore      |
|                 |              |         |      |         |                 | (upstream)                     |
+-----------------+--------------+---------+------+---------+-----------------+--------------------------------+
```

> The `--use-consul` flag tells glooctl to write configuration to Consul Key-Value storage

The route will send traffic from the root of the Gloo Gateway to the prefix `/api/pets` on the Pet Store application. You can test that by using `curl` against the Gateway Proxy URL:

```bash
curl <nomad-host>:8080/
```

If running on macOS or with Vagrant:

```bash
curl localhost:8080/
```

If running on Linux, use the Host IP on the `docker0` interface:

```bash
curl 172.17.0.1:8080/
```

Curl will return the following JSON payload from the Pet Store application.

```json
[{"id":1,"name":"Dog","status":"available"},{"id":2,"name":"Cat","status":"pending"}]
```

---

## Next Steps

Congratulations! You've successfully deployed Gloo to Nomad and created your first route. Now let's delve deeper into the world of [Gloo routing]({{< versioned_link_path fromRoot="/gloo_routing" >}}). 

Most of the existing tutorials for Gloo use Kubernetes as the underlying resource, but they can also use Nomad. Remember that all `glooctl` commands should be used with the `--use-consul` flag, and deployments will need to be orchestrated through Nomad instead of Kubernetes.
