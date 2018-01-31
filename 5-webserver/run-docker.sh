docker run                                               \
        --restart=always                                 \
        --hostname                                       \
        "ciscolive-service-content"                      \
        --name "ciscolive.service.content"               \
        -p "8888:80"                                     \
        -d                                               \
        ciscolive.service.content:latest       
exit $?