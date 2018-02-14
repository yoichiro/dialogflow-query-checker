# dialogflow-query-checker

This is a checker tool to validate a result phrase against each query sent to the Dialogflow.

To send a query to the Dialogflow, this uses the [Dialogflow Query API](https://dialogflow.com/docs/reference/agent/query).

## Install

You can get this tool from the following page:

https://github.com/yoichiro/dialogflow-query-checker/releases/latest

Download the archive file named `dialogflow-query-checker-[version].zip` and extract the file to some directory where you want to install. The archive file has some executable files for each OS and Architecture.

You can copy an executable file to a directory included into your `PATH` environment variable.

## Configuration File

To check your intents, actions and fulfillment, you need to create a configuration file. The format of the file is YAML. The structure of the file is:

```yaml
clientAccessToken: <CLIENT_ACCESS_TOKEN>
defaultLanguage: <DEFAULT_LANGUAGE>
tests:
  -
    condition:
      contexts:
        - <INPUT_CONTEXT>
      language: <LANGUAGE>
      query: <QUERY_STRING>
    expect:
      action: <ACTION_ID>
      intentName: <INTENT_ID>
      parameters:
        <PARAMETER_NAME>: <PARAMETER_VALUE>
      contexts:
        - <OUTPUT_CONTEXT>
      speech: <SPEECH_REGULAR_EXPRESSION>
```

* `CLIENT_ACCESS_TOKEN` - The client access token issued by the Dialogflow. You can get the token from the project configuration page of the your Dialogflow project.
* `DEFAULT_LANGUAGE` - This language is used, if the language value in each test definition is not specified.
* tests - This is an array which has each test case.
  * condition - This defines the condition of the query represented by contexts and a query.
    * `INPUT_CONTEXT` - The context ID when the query sends. You can specify multiple contexts, and also can omit.
    * `LANGUAGE` - The query language. The defaultLanguage is used when this value is omitted.
    * `QUERY_STRING` - The query string. This means "User says" in Dialogflow.
  * expect - This defines a expected result which should be returned from the Dialogflow.
    * `ACTION_ID` - The action ID determined by an intent.
    * `INTENT_ID` - The intent ID guessed by the query.
    * parameters - This defines parameters which were parsed from the query by the Dialogflow. 
      * `PARAMETER_NAME` - The parameter's name.
      * `PARAMETER_VALUE` - The parameter's value retrieved from the query phrase.
    * `OUTPUT_CONTEXT` - The context ID determined by the intent or the fulfillment. You can specify multiple contexts, and also can omit.
    * `SPEECH_REGULAR_EXPRESSION` - The regular expression to validate the response from the Dialogflow. 

In the `PARAMETER_VALUE` and the `SPEECH_REGULAR_EXPRESSION`, you can use macros. In the latest version, the following macros are supported:

* `${date.today}` - This is replaced to today's date string.
* `${date.tomorrow}` - This is replaced to tomorrow's date string.

The sample is like the following:

```yaml
clientAccessToken: ...
defaultLanguage: en
tests:
  -
    condition:
      contexts:
        - "input_condition"
      query: "How many times is a Google I/O in this year?" 
    expect:
      action: "event_info"
      intentName: "input.condition"
      parameters:
        event: "Google I/O"
        when: "2018"
      contexts:
        - "answered"
      speech: "^The event is the (1st)|(2nd)|(3rd)|([0-9]+th).$"
...
```

## Execute

After writing the configuration file, you can execute this tool to validate each query and the response based on the configuration file. To use this tool, execute the following command:

```bash
$ dialogflow-query-checker run <CONFIGURATION_FILE_PATH>
```

You see like the following output:

```bash
........F.....F........F.......
[input_condition How many times is a Google I/O in this year?] action is not same. expected:event_info actual:event_information
...
3 tests failed.
```

If all tests passed, exit status code is `0`. Otherwise, the code is `1`.

## For developers

You can build this tool by the following steps:

1. Install Go language environment.
1. `git clone git@github.com:yoichiro/dialogflow-query-checker.git`
1. `go get github.com/mitchellh/gox`
1. Change the version number written in the `verion.go` file.
1. Run `./cross_build.sh`
