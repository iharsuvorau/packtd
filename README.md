# packtd — Packt Publishing free e-book daemon 

PacktPublishing offers a free tech e-book each day on its website. The program below is able to display the last offer to the stdout or send an email to the specified recipient using the sender's credentials (gmail only).

## Usage

To see the last PacktPub offer just type in the terminal: 

`packtd`

To send the offer by email:

`packtd -to rec@example.com -from sender@host.com -pass 123`

You could automate the app using *launchctl* on OS X or *systemctl* on Linux.