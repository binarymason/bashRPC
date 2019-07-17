#!/usr/bin/env bash

set -e

test_secret=supersecret
cmd="bashRPC -c ./test/data/complex_config.yml"
url="localhost:8675"

setup() {
  go install

  echo "# starting bashRPC..."
  eval "$cmd" &>/dev/null &
  sleep 3
}

fail() {
  echo "!!! $*"
  echo "FAIL"
  exit 1
}

ok() {
  echo "--> OK"
}

cleanup() {
  pkill -f "$cmd"
}
trap cleanup EXIT

setup

assert_status_code() {
  local expected_code="$1"
  local status_code
  echo -e "\t* ${*:2}"
  status_code=$(curl -sw '%{http_code}' "${@:2}" | tail -n 1)
  if [ "$status_code" != "$expected_code" ]; then
    fail "expected status code to be $expected_code but got $status_code"
  fi

}
curl -f -H "Authorization: supersecret" localhost:8675

echo "# test 404 endpoint"
assert_status_code 404 -H "Authorization: $test_secret" "$url/foobar"
assert_status_code 404 -H "Authorization: $test_secret" "$url/404ftw"
assert_status_code 404 -H "Authorization: $test_secret" "$url/idontexist"
ok

echo "# test without authentication"
assert_status_code 401 "$url/version"
assert_status_code 401 "$url/multiline"
assert_status_code 401 "$url/git/status"
assert_status_code 401 "$url/fail"
assert_status_code 401 "$url/pipe"
assert_status_code 401 "$url/backslashes"


assert_response() {
  endpoint="$1"
  expected_response="$2"
  echo -e "\t* $endpoint"
  res=$(curl -H "Authorization: $test_secret" -s "$url$endpoint")
  if [ "$res" != "$expected_response" ]; then
    fail "expected response to be: $expected_response, but got: $res"
  fi
}

echo "# test version"
assert_response /version "1.2.3"
assert_response /nested/route "it works"
assert_response /fail "$(echo -e "womp womp\n: exit status 1")"
assert_response /pipe "piping works"
assert_response /backslashes "$(echo -e "it\nworks")"

multiline_response=$(cat <<'END_HEREDOC'
I could be a complex bash script
Today I am merely a bunch of strings
Indeed, I am a multiline string
I think, therefore I am.
END_HEREDOC
)

assert_response /multiline "$multiline_response"
ok
