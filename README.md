# Mamp Vhost Manager

![Main](https://github.com/capimichi/mamp-vhost-manager/raw/main/screenshot/main.png)

## The Objective

This project is a simple GUI for managing virtual hosts on a local MAMP installation.

It is born because I needed to manage a lot of virtual hosts on my local machine and I didn't want to do it manually.

It let you create, edit and delete virtual hosts.

## Status

It is still in development, but it is usable.

It currently works with apache2, but can be implemented to work with nginx.

It is born for OSX system, but can be implemented to work on other systems.

## What it does

It has a GUI that shows you all the virtual hosts you have on your machine.

It let you create, edit and delete virtual hosts.

It automatically creates the virtual host file in the apache2 folder, then adds the virtual host to the hosts file and let you restarts apache2.

It creates a new folder in MAMP config folder, so you can easily manage your virtual hosts.

It then adds an Include directive to the apache2.conf file, to include all hosts in the folder.

## What it doesn't do

It doesn't manage already existing virtual hosts.