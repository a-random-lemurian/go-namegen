# `namegen`

**namegen** is a small Go library for generating phrases from a JSON file, inspired by Endless Sky's phrase generation system.

## Changelog
* 0.2.0
  * The library will now return errors instead of dying with log.Fatalf().
  * Add Unmarshal() method to the library instead of requiring you to feed a file to the library. 

## Copyright

Some ship names for the default set were taken from the video game Endless Sky, under a GPL 3 license, particularly this [file](https://github.com/endless-sky/endless-sky/blob/master/data/human/names.txt).
