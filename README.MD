// copy protoc to GOPATH
https://github.com/google/protobuf/releases

// protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go

// goprotobuf lib
go get github.com/golang/protobuf/proto

// trans all pbfile to gofile
#!bin/sh
for file in ./*.proto
do
    if test -f $file
    then
        new_dir=${file%.*}
        mkdir $new_dir
        # echo ${new_dir}

        protoc -I=./ --go_out=./$new_dir/ $file
    fi
    if test -d $file
    then
        echo $file "is dir"
    fi
done
