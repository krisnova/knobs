# knobs
Kubernetes Native Open Broadcasting Software

_This is a work in progress. Pull requests accepted_

## Our Message

     The world is changing. 
     
     Which means in-person events aren't going to be what they used to be. 

     Right now there is a gap in the industry when it comes time to managing high quality digital events that are open, safe, fun, collaborative, and effective. 
     
     Pre-recorded sessions and proprietary UDP based zoom calls will never be able to replace the energy of an in person event.

     As an open source engineer, and Kubernetes expert I believe our community can come together to solve this problem in a healthy and powerful way. 

     This project is our work on doing just that.

        _Kris Nóva_

## Project components

The `knobs` project will be composed of the following components that can be composed together for live production and recording on Kubernetes like systems.

**Plugins**

These are small programs that define input and output mechanisms for knobs. These are things like Zoom streams, IRC bouncers, Slack bouncers, Video devices, HTML browsers, Images, Audio input.

Plugins can either be of type `source`, `output`, or `encoder` build around the files in [`libobs`](https://github.com/obsproject/obs-studio/blob/master/libobs/).

 - [obs-output.h](https://github.com/obsproject/obs-studio/blob/master/libobs/obs-output.h)
 - [obs-source.h](https://github.com/obsproject/obs-studio/blob/master/libobs/obs-source.h)
 - [obs-encoder.h](https://github.com/obsproject/obs-studio/blob/master/libobs/obs-encoder.h)

**Servers**

Servers speak the `rtmp` protocol as defined [here](https://wwwimages2.adobe.com/content/dam/acom/en/devnet/rtmp/pdf/rtmp_specification_1.0.pdf) and should listen by default on TCP port 1935.

Servers should both consume and broadcast data on the same TCP port. The RTMP protocol is very efficient and does not require very much overhead. Servers act as simple data transfer. 

**Clients**

Clients connect to servers using the `rtmp` protocol as defined [here](https://wwwimages2.adobe.com/content/dam/acom/en/devnet/rtmp/pdf/rtmp_specification_1.0.pdf) and will begin to stream their content.

## Subprojects

Below are subprojects that will need to be built out.

**CRDs**

We need to define an initial CRD representation of the header files describe in project components for `plugins`, `servers`, and `clients`.


**OBS Bindings**

We need to wrap the plugin header files in Go and provide a working example of a simple plugin that will work with projects like OBS studio.

We should also support rust because rust is cool.

## Philosophy

Every client will be responsible for managing their own outbound stream. Every client gets 1 stream and can use tools like OBS studio to compose their stream.

Real time communications is out of scope for the initial proposal. Solve this component independently using POTs or other real time communications channels. We hope to support seamless realtime communication that takes full advantage of UDP in a future release but this will need to be designed.
