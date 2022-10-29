## Azure のコンテナレジストリへのプッシュ方法

az login
az acr login --name ${レジストリ名}
docker build . -t ${ログインサーバ名}.azurecr.io/${リポジトリ名}
docker push ${ログインサーバ名}.azurecr.io/${リポジトリ名}
