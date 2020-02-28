FROM nginx:latest

COPY entrypoint.sh /entrypoint.sh
COPY templates /templates
RUN chmod +x /entrypoint.sh
ENTRYPOINT /entrypoint.sh
