'use strict';
 
var gulp = require('gulp'),
    sass = require('gulp-sass'),
    cssnano = require('gulp-cssnano'),
    sassModuleImporter = require('sass-module-importer'),
    rename = require('gulp-rename'),
    debug = require('gulp-debug'),
    autoprefixer = require('gulp-autoprefixer'),
    sourcemaps = require('gulp-sourcemaps');

var sassFiles = ['./src/assets/sass/**/*.scss'];

var config = {
  debug: true
};
 
gulp.task('sass', function () {
  return gulp.src(sassFiles)
    .pipe(sourcemaps.init())
    .pipe(debug({ title: 'sass files:' }))
    .pipe(sass({
        importer: sassModuleImporter()
    }).on('error', sass.logError))
    .pipe(debug({ title: 'sass-compiled-files:' }))
    .pipe(autoprefixer())
    .pipe(gulp.dest('./static/css'))
    .pipe(cssnano({ zindex: false }))
    .pipe(rename({ suffix: '.min' }))
    .pipe(debug({ title: 'minified-css-files:' }))
    .pipe(sourcemaps.write('./'))
    .pipe(gulp.dest('./static/css'));
});

gulp.task('default', ['sass']);
 
gulp.task('sass:watch', function () {
  gulp.watch(sassFiles, ['sass']);
});

gulp.task('watch', ['sass:watch']);