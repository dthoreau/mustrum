# mustrum
RESTful API for Raspberry PI LED HATs

Summary:
- Pimoroni ( shop.pimoroni.com ) make a series of LED arrays as HATs for the compact RaspBerry Pi single board computer.
- LEDSEQ ( https://ledseq.com/ ) make (or made) a LED display panel displaying animations and static images.
- A long running project I had needed a dynamically addressed panel.

This project is an attempt to re-write Pimoron's library https://github.com/pimoroni/unicorn-hat-hd, b
ut in Go, and bind it into a RESTful API that can be addressed programatically, and fed dynamic content.

A *possible* hardware design for this is presented by https://johnmccabe.net/technology/projects/ubercorn-gameframe-pt1/
