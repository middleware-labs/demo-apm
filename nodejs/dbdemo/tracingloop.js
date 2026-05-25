const http = require('http');
var request = require('request');


if (process.env.MW_AUTOGENERATE_TRACING_DATA) {
    setTimeout(() => {}, 5000);

    setInterval(()=>{

        let random = Math.random();

        if (random < 0.9) {
            http.get('http://localhost:3002/api/tutorials');
        }

        if (random < 0.1) {
            http.get('http://localhost:3002/error');
        }

        if (random < 0.3) {
            http.get('http://localhost:3002/500-error');
        }

        if (random < 0.5) {
            http.get('http://localhost:3002/504-error');
        }

        if (random < 0.6) {
            http.get('http://localhost:3002/304-error');
        }

        if (random < 0.7) {
            request.post('http://localhost:3002/api/tutorials', {
                "title": "git3",
                "description": "test description3"
            }, (error, response, body) => {
                let body_new = JSON.parse(body);
                request.put(`http://localhost:3002/api/tutorials/${body_new.id}`, {
                    "title": "git4",
                    "description": "test description4"
                    }, (puterror, putresponse, putbody) => {

                        request.delete(`http://localhost:3002/api/tutorials/${body_new.id}`);
                });
            });
        }

    },3000);

    // Database workload generator: exercises complex / locking / IO-bound /
    // write-churn queries so they surface in pg_stat_activity, each carrying a
    // traceparent SQL comment. dbload.js caps row count + VACUUMs on its own, so
    // this can run indefinitely without filling storage.
    setInterval(() => {
        let r = Math.random();
        if (r < 0.55)      http.get('http://localhost:3002/api/dbload/random');
        else if (r < 0.68) http.get('http://localhost:3002/api/dbload/complex');
        else if (r < 0.78) http.get('http://localhost:3002/api/dbload/blocking');
        else if (r < 0.86) http.get('http://localhost:3002/api/dbload/iowait');
        else if (r < 0.92) http.get('http://localhost:3002/api/dbload/insert');
        else if (r < 0.97) http.get('http://localhost:3002/api/dbload/seqscan');
        else               http.get('http://localhost:3002/api/dbload/delete');
    }, 4000);
}
