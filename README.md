# seichi-gateway-operator
宣言されたCRDに応じてseichi-gateway(bungeecord-proxy)のbackend-serverに関する設定を動的に制御するためのoperator

# memo

```sh
kubebuilder init --domain seichi.click --repo github.com/GiganticMinecraft/seichi-gateway-operator
```

```sh
kubebuilder create api --group seichiclick --version v1alpha1 --kind SeichiReviewGateway --controller --resource
```

imagePullPolicy: IfNotPresent -> config/crd/manager/manager.yml

```sh
kubebuilder create api --group seichiclick --version v1alpha1 --kind BungeeConfigTemplate --controller --resource
```
