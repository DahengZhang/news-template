FROM centos
COPY ./news /
COPY ./docs /docs/
COPY ./tpl /tpl/
EXPOSE 8080
CMD [ "/news" ]