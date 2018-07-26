/// <binding BeforeBuild='default' ProjectOpened='watch_all' />
var gulp = require('gulp');
var exec = require('gulp-exec');
var clean = require('gulp-clean');

gulp.task('compile_thrift', function () {
    return gulp.src('./Thrift/*.thrift')
        .pipe(exec('thrift -gen csharp -out ./Thrift <%= file.path %>'));
});

gulp.task('watch_all', function () {
    gulp.watch('./Thrift/*.thrift', ['compile_thrift']);
});

gulp.task('clean_thrift', function () {
    gulp.src('./Thrift/Data', { read: false })
        .pipe(clean());
});

gulp.task('thrift', ['clean_thrift', 'compile_thrift']);

gulp.task('clean_all', ['clean_thrift']);

gulp.task('default', ['clean_all', 'thrift']);