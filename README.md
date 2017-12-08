## docker 
- docker run -it --rm --name ma-instance -p 8080:8080 -v $GOPATH/src/beego_one_app:$GOPATH/src/beego_one_app -w $GOAPTH/src/beego_one_app ma-image

## create a new app
- bee new app_name
- bee version

## build
- go build -o bee_go main.go

## swagger
- bee generate docs

## generate 
- bee generate model  users -fields=name:string,email:string
- bee generate appcode -driver=mysql -conn="root:12345678@tcp(127.0.0.1:3306)/test" -level=1 //1:models; 2: models and controllers; 3: models and controllers and routes