/*
 * @Description:
 * @Author: dotwoo@gmail.com
 * @Github: https://github.com/dotwoo
 * @Date: 2021-07-22 13:10:06
 * @FilePath: /joy4/cgo/ffmpeg/ffmpeg.h
 */

#include <libavcodec/avcodec.h>
#include <libavformat/avformat.h>
#include <libavutil/avutil.h>
#include <libavutil/log.h>
#include <libavutil/opt.h>
#include <libswscale/swscale.h>
#include <string.h>

typedef struct {
  AVCodec *codec;
  AVCodecContext *codecCtx;
  AVFrame *frame;
  AVDictionary *options;
  int profile;
} FFCtx;

static inline int avcodec_profile_name_to_int(AVCodec *codec,
                                              const char *name) {
  const AVProfile *p;
  for (p = codec->profiles; p != NULL && p->profile != FF_PROFILE_UNKNOWN; p++)
    if (!strcasecmp(p->name, name)) return p->profile;
  return FF_PROFILE_UNKNOWN;
}
