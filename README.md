# powbot-launcher

A simple Go application that does the following:
* Check if the powbot directory has a JRE in it 
    * If not, it downloads the appropriate JRE for your platform from AdoptOpenJDK
* Check if the powbot directory has the latest client in it
    * Again, if not, it downloads the latest client from the website
* Bootstraps the client process using either cmd or bash

And while it does all this it prints pretty coloured text to a command window