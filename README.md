# Readme for the trackyourself project -- Firefox (ONLY, for now)

track yourself is a project that runs locally, is open source, and is built only to allow people to collect data about their web browsing, without any third party. Ideally, the project would also allow you to use other local data analysis tools to explore your data. 

## The project goal is to allow a user to track themselves using an open source tool, with no third party.

In order to do this the project has to use two parts: the tracker browser plugin and the database (and potentially an associated local server displaying the data).

### The Browser Plugin

This is supposed to be a simple plugin that tracks your data as much as possible, and send it all to your local server. The code is open source, so you can go see that the plugin is sending the data only to your local host. The plugin will create json of your data that it collects, and then will send it to you webserver/database that will also be running.

The recommended way to install this plugin would be to download the repo and add the plugin manually. for Firefox see:

 https://support.mozilla.org/en-US/kb/unable-install-add-ons-extensions-or-themes?redirectlocale=en-US&redirectslug=Unable+to+install+add-ons#w_you-are-asked-to-download-the-add-on-rather-than-installing-it

note that the manifest.json file you will want to point Firefox to is

```
trackyourself/plugin
```

You will also need to set up the local server in order for this to do anything.

### The local server - using docker-compose

This is a docker-compose setup that will have to be run manually. This gives you the ability to re-use the data however you may please!

In order to use this, you must have both Docker, and docker-compose installed. (See:
	https://docs.docker.com/install/linux/docker-ce/ubuntu/
	https://docs.docker.com/compose/install/)

To get the server running (you will have to restart the server on re-boot) cd into:

```
trackyourself/server
```

and build the images (if you have not since the last time you cloned the repo)

```
docker-compose build
```

and then finally, run the docker-compose setup with:

```
docker-compose up
```

Your server should be running, you can monitor that the server is receiving the requests through the log in the terminal.

### Other notes, where the project is going

This project is made to be simple, but is supposed to empower people to obtain their own data about themselves. Its hard to stop all the tracking (though there are ways) it is at least nice to see what could be put out there. I wanted to just give people a way to track some data for themselves.

I hope to add a frontend to this webserver that would also run locally. This would make it so that you could use a docker-compose setup to also run a webserver that you could view visualizations of you data on, WITH NO THIRD PARTY.

The ideal would be to allow the data to be limited enough so that a person could collect the data themselves. I would also love to let you chose what data to track about yourself.
