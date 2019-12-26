[Unit]
Description=Kubernetes Controller Manager
Documentation=https://github.com/kubernetes/kubernetes
 
[Service]
EnvironmentFile=-/opt/kubernetes/cfg/kube-controller-manager.conf
ExecStart=/opt/kubernetes/bin/kube-controller-manager \$KUBE_LOG \$KUBE_LOG_VERSION \\
\$KUBE_MASTER \$KUBE_ELECT $KUBE_ADDRESS \$KUBE_CLUSTER_IP_RANGE \$KUBE_CLUSTER_CIDR \\
\$KUBE_CLUSTER_NAME \$KUBE_CLUSTER_CERT_FILE \$KUBE_CLUSTER_KEY_FILE \$KUBE_ROOT_CA_FILE \\
\$KUCSTER_SA_PRIVATE_KEY $KUBE_ALLOCATE_NODE_CIDRS \$KUBE_HORIZONTAL_POD_AUTOSCALER_USE_REST_CLIENT
Restart=on-failure

[Install]
WantedBy=multi-user.target