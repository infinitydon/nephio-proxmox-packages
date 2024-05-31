<table style="border-collapse: collapse; border: none;">
  <tr style="border-collapse: collapse; border: none;">
    <td style="border-collapse: collapse; border: none;">
      <a href="http://www.openairinterface.org/">
         <img src="./docs/images/oai_final_logo.png" alt="" border=3 height=50 width=150>
         </img>
      </a>
    </td>
    <td style="border-collapse: collapse; border: none; vertical-align: center;">
      <b><font size = "5">OpenAirInterface 5G Core Network Function Topology for Nephio</font></b>
    </td>
  </tr>
</table>

[[_TOC_]]

# 1. Introduction

This repository contains Kpt packages and package variants for OAI operators and network functions. Designed for [Nephio](https://nephio.org/) release 2. 

**NOTICE**

All the files are published under BSD 3-Clause License and the yaml definations provided by nephio.

Git repository for OAI Operators --> `https://gitlab.eurecom.fr/development/oai-operators` 

Repository structure is below

```
.
├── database
├── docs
│   └── images
├── oai-amf
├── oai-ausf
├── oai-nrf
├── oai-nr-ue
│   └── nrue
├── oai-cp-operators
│   ├── crds
│   └── operator
├── oai-up-operators
│   ├── crds
│   └── operator
├── oai-repository
├── oai-rfsim-gnb
│   └── gnb
├── oai-smf
├── oai-udm
├── oai-udr
├── oai-upf-edge
├── package-variant
└── package-variant	Package variant for all nfs, database and operator
```

The package-variants are designed to deploy Nephio R2 deployment. 

# 2. How to deploy the toplogy on Nephio

## Step 0: Prerequisite

1. Make sure that you have a running core, core and edge cluster topology as defined in [nephio exercise](https://github.com/nephio-project/docs/blob/v1.0.1/user-guide/exercises.md). 

```bash
git clone https://github.com/OPENAIRINTERFACE/oai-packages.git
# IF NEEDED
git checkout <branch>
cd oai-packages
```

## Step 1 (Optional): Adding oai-packages repository in nephio repo list


```bash
kpt live init oai-repository
kpt live apply oai-repository
```

<details>
<summary>The output is similar to:</summary>

```console
installing inventory ResourceGroup CRD.
inventory update started
inventory update finished
apply phase started
repository.config.porch.kpt.dev/oai-packages apply successful
apply phase finished
reconcile phase started
repository.config.porch.kpt.dev/oai-packages reconcile pending
repository.config.porch.kpt.dev/oai-packages reconcile successful
reconcile phase finished
inventory update started
inventory update finished
apply result: 1 attempted, 1 successful, 0 skipped, 0 failed
reconcile result: 1 attempted, 1 successful, 0 skipped, 0 failed, 0 timed out
```
</details>

To verify the repository is on-boarded

```bash
kubectl get repositories | grep oai-packages
```

<details>
<summary>The output is similar to:</summary>

```console
NAME                      TYPE   CONTENT   DEPLOYMENT   READY   ADDRESS
oai-packages              git    Package   false        True    https://github.com/OPENAIRINTERFACE/oai-packages
```
</details>


## Step 2: Deploy core network functions package variant

Now start deploying the core network function package variant starting with database. 

1. UDR requires that the database is running 
2. All AMF, SMF, UDR, UDM, AUSF will wait for NRF container to be healthy

```bash
kubectl apply -f package-variant/database.yaml
```
<details>
<summary>The output is similar to:</summary>

```console
packagevariantset.config.porch.kpt.dev/core-oai-database created
```
</details>

After couple of seconds you will see database pod in core cluster in `oai-core` namespace

```bash
kubectl get pods -n oai-core --context core-admin@core
```
<details>
<summary>The output is similar to:</summary>

```console
NAME                     READY   STATUS    RESTARTS   AGE
mysql-5c6cb749bc-nsdsp   1/1     Running   0          47s
```
</details>

### Step 2.1 Deploy operators on the core and edge cluster

```bash
kubectl apply -f package-variant/operators-cp.yaml
kubectl apply -f package-variant/operators-up.yaml
```
<details>
<summary>The output is similar to:</summary>

```console
packagevariant.config.porch.kpt.dev/oai-cp-operators created
packagevariant.config.porch.kpt.dev/oai-up-operators created
```
</details>

After couple of mins oai-operator package will be in `main` branch of gitea Core and Edge repository. Core network control plane operators will be in Core cluster and user plane in Edge cluster.

```bash
kubectl get pods -n oai-cn-operators --context core-admin@core
kubectl get pods -n oai-cn-operators --context edge-admin@edge
```
<details>
<summary>The output is similar to:</summary>

```console
## control plane
NAME                                   READY   STATUS    RESTARTS   AGE
oai-amf-controller-55dfbf8c4-9qdl4     1/1     Running   0          2m24s
oai-ausf-controller-769d64999f-28ntm   1/1     Running   0          2m24s
oai-nrf-controller-67f556bf75-8svd5    1/1     Running   0          2m24s
oai-smf-controller-5b6db9f5cb-klfsw    1/1     Running   0          2m24s
oai-udm-controller-867847d4cb-qdrzl    1/1     Running   0          2m24s
oai-udr-controller-764f4bfdb9-zw622    1/1     Running   0          2m24s
## user plane
NAME                                  READY   STATUS    RESTARTS   AGE
oai-upf-controller-75cbc869cb-zchjl   1/1     Running   0          11s
```
</details>


### Step 2.3 Deploy Control Plane and User Plane Functions

Deploy control plane network functions AMF, SMF, NRF, UDR, UDM, AUSF

```bash
kubectl apply -f package-variant/nrf.yaml
kubectl apply -f package-variant/udm.yaml
kubectl apply -f package-variant/udr.yaml
kubectl apply -f package-variant/ausf.yaml
kubectl apply -f package-variant/amf.yaml
kubectl apply -f package-variant/smf.yaml
kubectl apply -f package-variant/upf-edge.yaml
```

In around 6-7 mins you will see all the control plane NFs in `oai-core` namespace in core cluster.

```bash
kubectl get pods -n oai-core --context core-admin@core
```
<details>
<summary>The output is similar to:</summary>

```console
NAME                             READY   STATUS    RESTARTS      AGE
amf-core-5667d55644-nkthg    1/1     Running   0             85s
ausf-core-77867547bb-vl92j   1/1     Running   0             85s
mysql-5c6cb749bc-hn26d       1/1     Running   0             15m
nrf-core-7c79d988f5-lszwk    1/1     Running   0             85s
smf-core-5966dfd454-fc484    1/1     Running   0             82s
udm-core-56f78c9c7c-44556    1/1     Running   0             85s
udr-core-6f685c97db-2vrb7    1/1     Running   0             85s
```
</details>

In few mins you will see upf instances in `oai-core` namespace in edge cluster respectively. 

```bash
kubectl get pods -n oai-core --context edge-admin@edge
```

<details>
<summary>The output is similar to:</summary>

```console
NAME                          READY   STATUS    RESTARTS   AGE
upf-edge-696976df64-gwn42   1/1     Running   0          42m
```
</details>

## Step 3: Check if the PFCP session between UPF and SMF is established

It is really important that the PFCP session is established between SMF and UPF. If there is no PFCP session then there is no point in moving forward. To check the session you have to read the logs of SMF or UPF. 

```bash
kubectl logs -n oai-core <edge-upf-pod-name> --context edge-admin@edge | grep 'Received SX HEARTBEAT REQUEST' | wc -l
```

<details>
<summary>The output is similar to (any value more than one means that the session is established successfully):</summary>

```console
26
```
</details>

In case you don't see a session the mostly probably it is a networking issue in the setup and UPF is not able to reach the SMF n4 ip-address. To check this we suggest that you go inside the SMF pod and install `tcpdump` and ping `n4` ip-address of UPF. 

To deploy the RAN network functions you can follow the steps from the Nephio documentation. 

# Contribution requests

In a general way, anybody who is willing can contribute on any part of the code in any network component.

Contributions can be simple bugfixes, advices and remarks on the design, architecture, coding/implementation.