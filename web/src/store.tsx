import {
  component$,
  createContext,
  Slot,
  useContextProvider,
  useStore,
  useTask$,
} from "@builder.io/qwik";
import { postData } from "~/utils/fetch";
import { NormalizeUrl, ValidateUrl } from "~/utils/url";
import type { SpeakLocale } from "qwik-speak";

export const LinkContext = createContext("link");
export const AlterContext = createContext("alter");

export interface IAlterContext {
  type: "success" | "error" | "warning" | "info";
  message: string;
  show: boolean;
  time: number;
}

export interface ILinkContext {
  rawUrl: string;
  ifLoading: boolean;
  shortUrl: string;
  ifError: boolean;
  ifShowResult: boolean;
}

export const clearState = (state: ILinkContext) => {
  state.shortUrl = "";
  state.ifError = false;
  state.ifShowResult = false;
};

const baseUrl = import.meta.env.VITE_API_DOMAIN;

const reqShort = (link: string) => {
  return postData(baseUrl + "/v1/slinks", {
    link,
  }).then((data) => {
    return data;
  });
};

export const composeShortUrl = (shortKey: string) => {
  return baseUrl + "/s/" + shortKey;
};

export const handleShort = async (
  state: ILinkContext,
  alterState: IAlterContext,
  locale: SpeakLocale
) => {
  const pass = ValidateUrl(state.rawUrl);
  const ifCn = locale.lang === "zh-CN";
  if (!pass) {
    state.ifError = true;
    showTip(
      alterState,
      "warning",
      ifCn
        ? "不要胡闹, 请输入正确的链接"
        : "Don't mess around, please enter the correct link",
      5000
    );
    return;
  }

  const url = NormalizeUrl(state.rawUrl);
  const res = await reqShort(url);
  console.log(res);
  if (res.code !== undefined) {
    state.ifError = true;
    showTip(
      alterState,
      "warning",
      ifCn
        ? "你很调皮, 这个链接不支持"
        : "You are naughty, this link is not supported",
      5000
    );
    return;
  }
  state.ifShowResult = true;
  state.shortUrl = composeShortUrl(res.key);
  showTip(
    alterState,
    "success",
    ifCn ? "很开心, 转换成功啦" : "Very happy, the conversion is successful",
    3000
  );
};

export const showTip = (
  state: IAlterContext,
  type: IAlterContext["type"],
  message: string,
  time = 1000
) => {
  state.type = type;
  state.message = message;
  state.show = true;
  state.time = time;
};

export const InputStore = component$(() => {
  const store = useStore<ILinkContext>({
    rawUrl: "",
    ifLoading: false,
    shortUrl: "https://short.cn/u/22disd",
    ifError: false,
    ifShowResult: false,
  });
  useTask$(async (ctx) => {
    ctx.track(() => store.rawUrl);
  });
  useContextProvider(LinkContext, store);

  const alterStore = useStore<IAlterContext>({
    type: "success",
    message: "",
    show: false,
    time: 1000,
  });
  useContextProvider(AlterContext, alterStore);
  return (
    <>
      <Slot />
    </>
  );
});
