db = connect( 'mongodb://root:example@127.0.0.1:27017/strava?authSource=admin' );

let result = db.workout.aggregate([{
    $match: {
        type: 'Run'
    }
}, {
    $set: {
        date: {
            $dateFromString: {
                dateString: '$start_date'
            }
        }
    }
}, {
    $group: {
        _id: {
            $dateToString: {
                format: '%Y-%m',
                date: '$date'
            }
        },
        totalMonthDistance: {
            $sum: {
                $divide: [
                    '$distance',
                    1000
                ]
            }
        }
    }
}, {
    $match: {
        totalMonthDistance: {
            $gte: 150
        }
    }
}]);

console.log(result);