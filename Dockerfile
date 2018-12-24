FROM    vikings/alpine:latest
LABEL   maintainer=ztao8607@gmail.com
COPY    bin/nsq-exporter /nsq-exporter
EXPOSE  80
ENTRYPOINT ["/nsq-exporter"]