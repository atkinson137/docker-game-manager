Remove-Item "vendor/tutorial/**/*.*"
Remove-Item "vendor/shared/**/*.*"
thrift --gen go -out vendor -r -I thrift thrift/tutorial.thrift