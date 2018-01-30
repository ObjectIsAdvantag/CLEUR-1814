docker run --network=isolated_broker_nw                  \
        --restart=always                                 \
        --hostname                                       \
        "ciscolive-service-content"                      \
        --name "ciscolive.service.content"               \
        -p "9876:8080"                                   \
        -d                                               \
        ciscolive.service.content:latest       
exit $?