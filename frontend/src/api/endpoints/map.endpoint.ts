import { GenericAbortSignal } from "axios";
import { MapDto } from "../models/map.model";
import api from "../utils/api";

export namespace MapEndpoint {
  export const getProperty = (
    latitude: number,
    longitude: number,
    radius: number
    // signal: GenericAbortSignal
  ): MapDto.Property[] => {
    console.log(latitude, longitude, radius);
    return mock;
  };
  // api.getAll("/consumers/q", MapDto.Property.array(), {
  //   params: { latitude, longitude, radius }
  //   // signal
  // });
}

const mock = [
  {
    balanceHolder: "ПАО «МОЭК»",
    consumerAddress: {
      unom: 1244,
      address:
        "Российская Федерация, город Москва, внутригородская территория муниципальный округ Гольяново, Байкальская улица, дом 37",
      municipalDistrict: "муниципальный округ Гольяново",
      border: [
        {
          Key: "coordinates",
          Value: [
            [
              [37.814842591, 55.81773552],
              [37.814852144, 55.817634451],
              [37.814057571, 55.817617621],
              [37.814051928, 55.817709249],
              [37.814077452, 55.817709183],
              [37.814074565, 55.817746914],
              [37.81407091, 55.817788239],
              [37.81406872, 55.817813393],
              [37.814186796, 55.817816685],
              [37.814183982, 55.817863396],
              [37.814139308, 55.817862612],
              [37.814137205, 55.817898544],
              [37.814181878, 55.817899328],
              [37.814181197, 55.8179137],
              [37.814175655, 55.818017902],
              [37.814402213, 55.818021365],
              [37.814411995, 55.817849791],
              [37.814461447, 55.817849663],
              [37.814461222, 55.817821821],
              [37.814700586, 55.817831088],
              [37.814703675, 55.817818506],
              [37.814834499, 55.817819968],
              [37.814917451, 55.817819755],
              [37.814926351, 55.817736652],
              [37.814842591, 55.81773552]
            ]
          ]
        },
        {
          Key: "type",
          Value: "Polygon"
        }
      ],
      center: {
        type: "Point",
        coordinates: [37.814445085, 55.817819131]
      }
    },
    Source: "ТЭЦ №23",
    number: "04-06-0601/010",
    type: "ЦТП",
    locationType: "Отдельно стоящий",
    district: "район Гольяново",
    commissioningDate: -126230400000,
    heatingPointAddress: {
      unom: 2402445,
      address:
        "Российская Федерация, город Москва, внутригородская территория муниципальный округ Гольяново, Байкальская улица, дом 43, строение 1",
      municipalDistrict: "муниципальный округ Гольяново",
      border: [
        {
          Key: "coordinates",
          Value: [
            [
              [37.818884555, 55.818366397],
              [37.819012198, 55.81836876],
              [37.819023297, 55.818165745],
              [37.818952309, 55.818164392],
              [37.818908834, 55.818163564],
              [37.818898848, 55.818163373],
              [37.818895656, 55.818163381],
              [37.818884555, 55.818366397]
            ]
          ]
        },
        {
          Key: "type",
          Value: "Polygon"
        }
      ],
      center: {
        type: "Point",
        coordinates: [37.818953926, 55.818266071]
      }
    }
  },
  {
    balanceHolder: "ПАО «МОЭК»",
    consumerAddress: {
      unom: 7381,
      address:
        "Российская Федерация, город Москва, внутригородская территория муниципальный округ Сокольники, Егерская улица, дом 10",
      municipalDistrict: "муниципальный округ Сокольники",
      border: [
        {
          Key: "coordinates",
          Value: [
            [
              [37.684773727, 55.793951496],
              [37.684520015, 55.794034572],
              [37.684670528, 55.79418064],
              [37.68492424, 55.794097565],
              [37.684773727, 55.793951496]
            ]
          ]
        },
        {
          Key: "type",
          Value: "Polygon"
        }
      ],
      center: {
        type: "Point",
        coordinates: [37.684722127, 55.794066068]
      }
    },
    Source: "ТЭЦ №23",
    number: "04-02-0614/021",
    type: "ЦТП",
    locationType: "Отдельно стоящий",
    district: "район Сокольники",
    commissioningDate: 63072000000,
    heatingPointAddress: {
      unom: 2400806,
      address:
        "Российская Федерация, город Москва, внутригородская территория муниципальный округ Сокольники, Охотничья улица, дом 6, строение 2",
      municipalDistrict: "муниципальный округ Сокольники",
      border: [
        {
          Key: "coordinates",
          Value: [
            [
              [37.686373961, 55.794731291],
              [37.686197486, 55.794667398],
              [37.68608156, 55.79476909],
              [37.686258017, 55.794832992],
              [37.686373961, 55.794731291]
            ]
          ]
        },
        {
          Key: "type",
          Value: "Polygon"
        }
      ],
      center: {
        type: "Point",
        coordinates: [37.68622776, 55.79475019]
      }
    }
  }
];