# swap
Swap is a tool for change the default profile in your ~/.aws/credentials and ~/.aws/config files.

## Why use this tool
Swap allows you to change the default aws profile you are currently using. Besides this, it allows you to make configurations or change your ~/.aws/credentials file and save them or restore the ones you have already saved.

## Install Swap
Install the latest version:

``go install github.com/jordabrell/swap@latest ``

## How to use it
First of all we are going to save our configuration with the following command:

``swap save``

Once our configuration is saved, we could switch between our aws profiles using this command:

``swap profile <profile>``

When the profile has been changed, if we want to return to our initial configuration, we can do it with the following command:

``swap restore``
