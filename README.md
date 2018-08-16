# CodeMAC

## About the project

The philosophy behind this tool is know about the code.

As a developer, we have been reading the code and trying to understand it.But why can't there be a tool/helper that converts the code into plain/simple language and explain about the semantics of the code. Personally I thought of having such a handy companion many times.

The idea of CodeMAC is Measure , Analyze and Convert the code and so named this tool as CodeMAC .

## What does CodeMAC tool does?

CodeMAC does three things(as an initial plan).

1.Metrics  : Measures the code
2.Analyzer : Analyzes the code based on the metrics.
3.Converter: Converter converts the code into simple understandable language.

## How does it work ?

CodeMAC tool, parses through the code and identify which language/scripting the code has been written.It looks for specific language add-on /rule engine and start writing the code in .MD [MarkDown] file.

## Which languages/scripting does it support ?

The initial version, we are planing for Go(GoLang) source codes.We will have to work on other languages/scripting at a later stage based upon contributions.

## Are there any specifications/standards that it follows ?

At this point of time no.But probably in future.

## Is it open source ?

Yes, this is an open source project, will be published on GitHub.Com as a public repository with MIT license(May be changed to Apache later).

## Example of intended output

For a given source code file

#### 1.Metrics

```
Source code language:Go
File Name: sample.go
File Size: 2,722 bytes

Total Lines : 83

Total Functions: 12

Total Procedures: 

Total Methods: 6

Total Global Variables: 11

Total Local Variabls: 6

Total References: 3

Total Loops : 5

Total Conditions : 23
```
#### 2.Analyzer

```
func Wrapper : 3 nested loops are written.

func Wrapper : 12 conditions are written.
```
#### 3.Convertor

```
func WrapClient(w ...client.Wrapper) Option {
    return func(o *Options) {
        // apply in reverse
        for i := len(w); i > 0; i-- {
            o.Client = w[i-1](o.Client)
        }
    }
The above code will be converted as the following textual format[Not exactly but yes like below at this point of time].

-- WrapClient is a function with w is variable parameters of type client.Wrapper and returns Option

-- function implementation starts

-- returns an anonymus function with o *Option

-- Loop from i to len(w) greater than 0.Loop decreased 1.

-- *Options object is assigned to slice of w[i-1](o.Client)

-- for Loop ends.

-- function ends.

```

## Who is initial contributor or developer for this project ?

At this point of time Jiten.I am looking for more contributors.

