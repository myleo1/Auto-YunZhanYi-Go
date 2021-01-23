FROM alpine:latest
LABEL maintainer="leo <leo@leom.me>" \
	version="v0.1" \
	description="AUTO-YunZhanYi-Go-ZJGSU"
WORKDIR /root
ADD Auto-YunZhanYi-Go init.sh /home/
RUN echo '04       23      *       *       *       /home/Auto-YunZhanYi-Go' > /etc/crontabs/root \
	&& chmod -R 777 /home
CMD ["/bin/sh","/home/init.sh"]