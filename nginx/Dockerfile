FROM ubuntu:14.04
MAINTAINER Tyler Geery

# Install latest nginx
RUN apt-get update
RUN sudo apt-get -y install nginx

# Remove the default Nginx configuration file
RUN rm -v /etc/nginx/nginx.conf

# Copy a configuration file from the current directory
ADD nginx/nginx.conf /etc/nginx/

# Append "daemon off;" to the configuration file
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

# Expose port
EXPOSE 80

# Set the default command to execute when creating a new container
CMD service nginx start
