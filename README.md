# kubectl-gpt
kubectl-gpt is kubectl plus that is increased intelliSense and code generate ability by chatgpt.
- Quick switch cluster,that it loaded directory  k8s cluster config files.
- Enhance kubectl command intelliSense. Such as get resources  ...
- Support chatGPT generate k8s resources yaml and kubectl commands.

# install 
```bash
go install github.com/yoyofx/kubectl-gpt@v1.0.1
```

# Installation location:
$GOPATH

Add $GOPATH to $PATH Environment variable
```bash
export GOROOT=/usr/lib/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

# ENV Setting 
*KUBE_CHATGPT_CONFIG* Used to kubectl-gpt switch clusters, that it loaded directory  k8s cluster config files.

*KUBE_CHATGPT_TOKEN* It is chatGPT's token . It used to query chatGPT k8s resource yaml and kubectl commands shell.

```bash
KUBE_CHATGPT_CONFIG=/data/clusters/
KUBE_CHATGPT_TOKEN=sk-xxxxxxxeeei1nD3uKxwwe332efwfexxxffss
```
# Demo
![image](https://user-images.githubusercontent.com/4504853/224553454-7342794f-dd23-4da9-ac27-ee24fadaa3ab.png)
![image](https://user-images.githubusercontent.com/4504853/224553488-dccd3343-2ad8-422f-aaf6-d8f99c3dee90.png)
![image](https://user-images.githubusercontent.com/4504853/224553565-364cee18-d939-405f-a47a-90bc6c0e24a6.png)
## Demo2 , chatGPT generate k8s resource yaml 
![image](https://user-images.githubusercontent.com/4504853/224553662-80afe3fd-d346-40f2-b561-c242f18f9a42.png)
## Demo3 , chatGPT generate kubectl command , that is get deploy on all namespace.
![image](https://user-images.githubusercontent.com/4504853/224554075-3498c4eb-d408-4530-a11f-9b9698491683.png)
## Demo4 , quick switch cluster
![image](https://user-images.githubusercontent.com/4504853/224554180-3034dd9a-adef-4be0-9a78-59b1e012ebcb.png)

