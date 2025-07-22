#!/bin/bash

NAMESPACE="<some_namespace>"
POD_PREFIX="<some_pod_prefix>"
DATE_TAG="<some_date>"
LOG_PREFIX="log"

# Get list of matching pods
PODS=$(kubectl get pods -n "$NAMESPACE" --no-headers -o custom-columns=":metadata.name" | grep "^$POD_PREFIX")

# Check if any pods were found
if [ -z "$PODS" ]; then
  echo "No pods found with prefix '$POD_PREFIX' in namespace '$NAMESPACE'"
  exit 1
fi

# Loop through each pod and get logs
INDEX=1
for POD in $PODS; do
  LOG_FILE="${LOG_PREFIX}_${DATE_TAG}_${INDEX}.log"
  echo "Fetching logs for pod: $POD -> $LOG_FILE"
  kubectl logs "$POD" -n "$NAMESPACE" > "$LOG_FILE" 2>&1
  ((INDEX++))
done
echo "Log collection complete."

