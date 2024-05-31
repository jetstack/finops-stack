#!/usr/bin/env python
from diagrams import Cluster, Diagram, Edge
import os
from diagrams.onprem.monitoring import Prometheus, Grafana
from diagrams.onprem.client import Users
from diagrams.custom import Custom

from diagrams.k8s.compute import Deploy
from diagrams.k8s.others import CRD
from diagrams.generic.place import Datacenter

with Diagram("Architecture",
             filename="../content/assets/architecture",
             outformat="png",
             show=False,
             direction="TB"):

    users = Users("Users")

    with Cluster("Cluster"):
        with Cluster("[ns] finops-stack"):
            prom = Prometheus("Prometheus")
            graf = Grafana("Grafana")
            kyverno = Custom("Kyverno", os.getcwd()+"/resources/kyverno.png")
            finops_policies = CRD("Finops Policies")
            opencost = Custom("Opencost", os.getcwd() +
                              "/resources/opencost.png")

            vpatron = Deploy("vpatron")
            officehours = Deploy("office hours")
            limitranger = Deploy("limit ranger")
            autospot = Deploy("autospot")

            # Hack to get more in the same cluster
            vpatron - Edge(color="transparent") - officehours
            limitranger - Edge(color="transparent") - autospot

        for i in range(0, 0):
            with Cluster(f"[ns] Application {i}"):
                Deploy(f"Application {i}")

    aws = Custom("AWS", os.getcwd()+'/resources/aws.png')
    azure = Custom("Azure", os.getcwd()+'/resources/azure.png')
    gcp = Custom("GCP", os.getcwd()+'/resources/google-cloud.png')
    onprem = Datacenter("On-Prem")

    users >> graf
    graf >> prom
    kyverno << Edge(style="dashed") << finops_policies
    prom >> opencost

    for i in [aws, azure, gcp, onprem]:
        opencost << Edge(label="Cost Data", style='dotted') << i
