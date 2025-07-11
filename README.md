# trans new tab page

https://diyhrt.market/api/

## development

1. Install [air](https://github.com/air-verse/air)
2. Run air

```sh
air dev
```

## Config and Profiles

This tool works with profiles. The default profile is `startpage`. If you want to load another profile just write it as command line arg after the command. To write a config File you can create the files here:

- `{profile}.toml`
- `.{profile}.toml`
- `~/.config/startpage/{profile}.toml`

## TODO

- implement ctl
- implement fetching in intervals
- actually building and figuring out how I should do that
- writing documentation
- implement autocomplete with a nice go backend and fast communication. Since it all runs locally nobody should have privacy concerns NEEDS TO BE ABLE TO TOGGLED OFF FOR DEMO PAGE
