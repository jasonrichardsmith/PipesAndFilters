FROM alpine
RUN apk update && apk upgrade
RUN apk add bash
EXPOSE 8082
ENV EXEC_SERVICE /test.sh 
ENV PORT 8082
ENTRYPOINT ["/http"]
ADD http/http /
ADD test.sh /
