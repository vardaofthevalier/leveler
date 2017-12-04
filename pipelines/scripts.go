package pipelines

var RunJobScript = `
#!/bin/bash

echo "Getting credentials from vault..."
printf %s\\n $(ls -1) /data/get_credentials/) | xargs -n 1 -P 0 -I {} /data/get_credentials/{} || exit 1

echo "Setting credentials..."
printf %s\\n $(ls -1) /data/set_credentials/) | xargs -n 1 -P 0 -I {} /data/set_credentials{} || exit 1

echo "Syncing inputs..."
printf %s\\n $(ls -1 /data/sync_inputs/) | xargs -n 1 -P 0 -I {} /data/sync_inputs/{} || exit 1

echo "Running job..."
/data/run.sh || exit 1

echo "Syncing outputs..."
printf %s\\n $(ls -1 /data/sync_outputs/) | xargs -n 1 -P 0 -I {} /data/sync_outputs/{} || exit 1

echo "Job finished successfully!"
`