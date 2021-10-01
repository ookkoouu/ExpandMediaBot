cd $(dirname $0)
if [ -z $GH_TOKEN ]; then
  echo "GH_TOKEN is not set."
  exit
fi
curl -sLJO -H 'Accept: application/octet-stream' \
  "https://$GH_TOKEN@api.github.com/repos/ookkoouu/ExpandMediaBot/releases/assets/$(\
    curl -sL https://$GH_TOKEN@api.github.com/repos/ookkoouu/ExpandMediaBot/releases/latest \
      |jq '.assets[] | select(.name | contains("arm")) | .id')"
