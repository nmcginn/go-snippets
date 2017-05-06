#!/usr/bin/env python

import boto3

if __name__ == '__main__':
    client = boto3.client('ecs')
    result = client.run_task(
        cluster='merica',
        taskDefinition='reverse',
        overrides={
            'containerOverrides': [
                {
                    'name': 'reverse',
                    'command': [
                        '/go-snippets'
                     ],
                    'environment': [
                        {
                            'name': 'INPUT',
                            'value': 'EDIRREVO HSARC'
                        }
                     ]
                }
            ]
        }
    )
    print result
