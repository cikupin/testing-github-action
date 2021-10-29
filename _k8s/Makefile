.PHONY: install
install:
	helm upgrade --install loki --namespace=loki-stack loki/loki-stack --set grafana.enabled=true
	@echo "Grafana admin password"
	@kubectl get secret --namespace <YOUR-NAMESPACE> loki-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
	kubectl apply -f 1_create_namespace.yaml
	kubectl apply -f 2_deployment.yaml
	kubectl apply -f 3_service.yaml
	@echo "[DONE] Don't forget to expose app service & grafana"

.PHONY: delete-app
delete-app:
	kubectl delete -f 3_service.yaml
	kubectl delete -f 2_deployment.yaml
	kubectl delete -f 1_create_namespace.yaml

.PHONY: delete-all
delete-all:
	helm uninstall loki -n loki-stack
	kl delete namespace loki-stack