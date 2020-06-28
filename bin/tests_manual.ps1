Set-Location terraform-provider-todo
go build
Set-Location ..
docker-compose build
docker-compose up -d
# Add something to import as a data source
curl.exe -i -X POST -H 'Content-Type: application/spkane.todo-list.v1+json' --% -d "{\"description\":\"go shopping\",\"completed\":false}" http://127.0.0.1:8080/
Set-Location terraform-tests
Remove-Item -ErrorAction SilentlyContinue -Force terraform-provider-todo
Copy-Item ../terraform-provider-todo/terraform-provider-todo.exe .
terraform init --get --upgrade=true
$env:TF_LOG = 1
terraform apply
curl.exe -i http://127.0.0.1:8080/
docker-compose down
Remove-Item -Force terraform-provider-todo.exe
