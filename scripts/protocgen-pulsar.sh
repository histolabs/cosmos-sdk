# this script is for generating protobuf files for the new google.golang.org/protobuf API
set -eo pipefail



echo "Generate x/tx"
(cd x/tx; make codegen)