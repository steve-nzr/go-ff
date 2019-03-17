#!/bin/bash

set -euo pipefail

helm install --name go-ff-storage stable/postgresql \
  --set postgresqlPassword=0MvpRaEaIOuUep4xF5lF72z0OWDSPFhlFDEIkKpXC1TnRC5aDQ12Rt5CjjUg \
  --set postgresqlDatabase=go-ff

helm install --name go-ff-cache stable/mysql \
  --set mysqlRootPassword=oGBqEtgpCClMrLZqpKjVpH5I2WnHi39TBut1tbcvTboSRIkfnQi32Zhw45Lv \
  --set mysqlDatabase=go-ff

helm install --name go-ff-amqp stable/rabbitmq \
  --set rabbitmq.password=pPvqWUumUfDX85AFHEkwEmDhDKz25hkjXEjKaCHUuWJcSUJbbTxs52zjBJbh
