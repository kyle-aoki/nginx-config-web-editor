ssh $1 "rm -rf nginx-config-web-editor"
GOOS=linux GOARCH=amd64 go build
scp nginx-config-web-editor $1:~/nginx-config-web-editor
rm nginx-config-web-editor
scp cfg.yaml $1:~/cfg.yaml
