# Birdie programming exercise

This exercise is a small programming problem for you to solve. The goal is that we can get an idea
of how you approach a programming task and to have a chat with you about how you solved it. Your
solution doesn't need to be perfect -- just implement it the way you'd normally work on a task and
make sure to send the solution to us before the technical interview.

Please read the description and the notes below carefully before you start. Good luck!

## Description

The task is based on a part of our backend at Birdie: we get data in JSON format from our customers
and in some cases we apply "transformations" to modify the data before storing it.

The goal of the exercise is to build a very simple version of that. We want a command-line program
called `transform` that loads JSON data from a file, transforms it according to its arguments, and
stores the output to another file.

For example, let's say we have a file `input.json` with this content:

    {
        "color": "blue",
        "number": "two",
        "pet": "cat"
    }

and we run the program like this:

    ./transform delete:color set:number:three rename:pet:animal

That should produce a file `output.json` with this content:

    {
        "number": "three",
        "animal": "cat"
    }


## Input and output

To keep it simple, the input is always in a file called `input.json` in the current directory. The
contents of the file is a JSON document, specifically a JSON object. A JSON object maps keys to
values. For example, this object:

    {"color": "blue", "number": "two"}

maps the key `color` to value `blue` and the key `number` to value `two`. The keys and values of
the object are all strings.

The program writes its output to a file called `output.json`, also in the current directory. It will
also contain a JSON object whose keys and values are strings.

## Transformations

Each command-line argument for the program is one transformation. The program applies those
transformations in the order that they're given on the command-line.

There are three types of transformation:

 * A `delete` transformation removes a key from the object. The argument `delete:key1` means to
   remove key `key1`. If the key doesn't exist in the object, then this transformation doesn't
   have any effect.
 * A `set` transformation sets the value for a key in the object. The argument `set:key1:val1` means
   to set key `key1` to value `val1`. If key `key1` already exists, its previous value is discarded.
 * A `rename` transformation changes the key of a value in the object. The argument
   `rename:key1:key2` means that the key `key1` is renamed to `key2` and the previous value (if any)
   of `key2` is discarded. However, if `key1` doesn't exist then the transformation has no effect.

Keys and values that aren't mentioned in the transformation are not changed. If the `transform`
program is run without any arguments, then the output is the same as the input.

You can assume that each key or value consists of one or more letters or numbers, both for
`input.json` and the keys/values given on the command-line. That means you don't need to handle
cases such as keys that contain a colon (`:`) or a space character (` `).

## Notes

 * The exercise should not take more than about 2 hours. If you're not finished in that time, just
   document which parts of the task are completed so we know what to look for.
 * We work mainly with Go, but if you're not comfortable with Go feel free to use another
   programming language.
 * You can use whatever tools or libraries you like for the implementation.
