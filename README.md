# JCOLLECTOR

## AIRFLOW
### Installation

install Airflow with the following command in k8s:

```bash
helm repo add apache-airflow https://airflow.apache.org
helm repo update
kubectl create namespace airflow
helmfile -e <environment> apply
```
