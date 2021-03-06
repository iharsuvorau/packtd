# packtd — Packt Publishing free e-book daemon 

PacktPublishing [offers a free tech e-book](https://www.packtpub.com/packt/offers/free-learning) each day on its website. The program below is able to display the last offer to the stdout or send an email to the specified recipient using the sender's credentials (gmail only).

## Usage

To see the last PacktPub offer just type in the terminal: 
```
$ packtd
```

To send the offer by email:

```
$ packtd -to recipient@example.com -from sender@host.com -pass 123
```

## Automation

You could automate the app using *launchctl* on OS X or *systemctl* on Linux.

### For mac

Install [terminal-notifier](https://github.com/julienXX/terminal-notifier/blob/master/README.markdown) to use system notifications.
```
$ brew install terminal-notifier
```

The bash script checks internet connection and runs our programs. Edit *launch_packtd.sh.example* specifying paths to binaries: `packtd` and `terminal-notifier`.

Put *launch_packtd.sh.example* at *~/Library/LaunchAgents/launch_packtd.sh*:
```
cp launch_packtd.sh.example ~/Library/LaunchAgents/launch_packtd.sh
```

Make it executable:
```
chmod +x ~/Library/LaunchAgents/launch_packtd.sh
``` 

Launchctl uses its own XML-based files for services. Edit *packtd.plist.example* specifying the path to the bash script file. 

Put *packtd.plist.example* at *~/Library/LaunchAgents/packtd.plist*:
```
cp packtd.plist.example ~/Library/LaunchAgents/packtd.plist
```

Load the daemon:
```
$ launchctl load ~/Library/LaunchAgents/packtd.plist
```

That's it. You will recieve daily notifications with free PacktPub's e-books. On click you can check the web page and grab an item.