local tcnt = 1
wrk.method = "POST"
wrk.headers['Content-Type'] = "application/json"
--wrk.body = '{"uid": 1,"activityId": 1,"goodsId": 1,"stockId": 1,"buyCnt": 1,"accessTime": 0}'

function setup(thread)
    thread:set("uid", tcnt * 1000)
    tcnt = tcnt + 1
end

request = function()
    local uid = wrk.thread:get("uid")
    local body = string.format(
        '{"uid": %d,"activityId": 1,"goodsId": 1,"stockId": 1,"buyCnt": 1,"accessTime": 0}',
        uid
    )
    uid = uid + 1
    wrk.thread:set("uid", uid)
    return wrk.format(wrk.method, wrk.path, wrk.headers, body)
end

response = function(status, headers, body)
    if status ~= 200 then
        io.write("------------------------------\n")
        io.write("Response with status: ".. status .."\n")
        io.write("[response] Body:\n")
        io.write(body .. "\n")
        io.write("------------------------------\n\n")
    end
end