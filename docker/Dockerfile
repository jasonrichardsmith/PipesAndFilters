FROM alpine
RUN apk update && apk upgrade
RUN apk add bash
EXPOSE 8080
ENV EXEC_SERVICE /test.sh 
ENV PORT 8080
ENTRYPOINT ["/http"]
ADD app/http /
ADD test.sh /
