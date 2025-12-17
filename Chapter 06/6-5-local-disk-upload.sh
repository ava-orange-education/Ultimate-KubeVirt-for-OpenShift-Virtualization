virtctl image-upload \
--uploadproxy-url=$(kubectl get service cdi-uploadproxy -n kubevirt-cdi -o jsonpath='{.spec.clusterIP}') \
--pvc-name=my-upload-pvc \
--pvc-namespace=default \
--image-path=/path/to/your/local/disk.img \
--insecure