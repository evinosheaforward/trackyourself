# Readme for trackyourself (v0.1) -- Firefox (ONLY, for now)

trackyourself is a project that will out-of-the-box completely run locally, is open source, and is built to allow people to collect data about their web browsing, without any third party. Ideally, the project would also allow you to use other local data analysis tools to explore your data.

## The project goal is to allow a user to track themselves using an open source tool, with no second party.

In order to do this the project has to use two parts: the tracker (browser plugin) and the recorder (a golang server that writes json to the database).

### The Browser Plugin

This is supposed to be a simple plugin that tracks your data and send it all to your local server. Right now, all it tracks is when you send an HTTP request with type mainframe. The code is open source, so you can go see that the plugin is sending the data only to your local host. The plugin will create json of your data that it collects, and then will send it to the webserver/database that will also be running.

The recommended way to install this plugin would be to to add the plugin as a temporary add-on. [See the MDN Docs](https://support.mozilla.org/en-US/kb/unable-install-add-ons-extensions-or-themes?redirectlocale=en-US&redirectslug=Unable+to+install+add-ons#w_you-are-asked-to-download-the-add-on-rather-than-installing-it)

First, clone the repository:

```
git clone https://github.com/evinosheaforward/trackyourself
```

and then add the plugin manually by going to about:debugging in your browser. The command line to do this is:

```
firefox about:debugging
```

Then click the button that says "Load Temporary Add-on...". Find the repository and select the manifest.json file:

```
trackyourself/plugin/manifest.json
```

You will also need to set up the local server in order to save the collected data.

### The local server - using docker-compose

This is a docker-compose setup that will have to be run manually. This gives you the ability to re-use the data however you may please!

In order to use this, you must have both Docker, and docker-compose installed. (See:
	[docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/) and [docker-compose](	https://docs.docker.com/compose/install/))

To get the server running (you will have to restart the server on re-boot) cd into:

```
trackyourself/server
```

and build the images (if you have not since the last time you cloned the repo) by running:

```
s/build
```

and then finally, run the docker-compose setup with:

```
s/run
```

Your server should be running, you can monitor that the server is receiving the requests through the log in the terminal.

### Ideas for things to do with the data.  
- local webserver displaying the data
- using the data in a jupyter notebook (NLP?)
- creating event-driven applications such as notifying you when you, e.g. spend a certain amount of time on a website.

### The Point of the recorder
- accept http requests of json
- insert the json into a column in the db

If you want to collect other data, you will be ale to use the recorder if you can locally send an http request of json to the recorder

### Other notes, where the project is going

This project is made to be simple, but is supposed to empower people to obtain their own data about themselves. Its hard to stop all the tracking (though there are ways) it is at least nice to see what could be put out there. I wanted to just give people a way to track some data for themselves.

I hope to add a frontend to the webserver that would also run locally. This would make it so that you could use a docker-compose setup to also run a webserver that you could view visualizations of your data on, WITHOUT A SECOND PARTY.

The ideal would be to allow the data to be limited enough so that a person could collect the data themselves. I would also love to let you chose what data to track about yourself, and also give people the opportunity to build further projects themselves that use the data.
