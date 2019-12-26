[Unit]
Description=Kubernetes Scheduler
Documentation=https://github.com/kubernetes/kubernetes

[Service]
EnvironmentFile=-/opt/kubernetes/cfg/kube-scheduler.conf
ExecStart=/opt/kubernetes/bin/kube-scheduler \$KUBE_LOG \$KUBE_LOG_VERSION \$KUBE_MASTER \$KUBE_ELETCT
Restart=on-failure

[Install]
WantedBy=multi-user.target